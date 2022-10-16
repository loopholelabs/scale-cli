/*
	Copyright 2022 Loophole Labs

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		   http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package auth

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog"
	"net"
	"strconv"
	"sync"
	"time"
)

const (
	Listen       = "127.0.0.1:8085"
	CallbackPath = "/auth/callback"
	RedirectURL  = "http://" + Listen + CallbackPath
)

type Auth struct {
	Username     string `json:"username"`
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type server struct {
	logger   zerolog.Logger
	app      *fiber.App
	listener net.Listener
	wg       sync.WaitGroup
	auth     chan *Auth
	err      chan error
}

func Do(logger zerolog.Logger) (*Auth, error) {
	s := &server{
		logger: logger.With().Str(zerolog.CallerFieldName, "AUTH").Logger(),
		app: fiber.New(fiber.Config{
			DisableStartupMessage: true,
			ReadTimeout:           time.Second * 10,
			WriteTimeout:          time.Second * 10,
			IdleTimeout:           time.Second * 10,
			JSONEncoder:           json.Marshal,
			JSONDecoder:           json.Unmarshal,
		}),
		auth: make(chan *Auth, 1),
		err:  make(chan error, 1),
	}

	s.init()

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		err := s.start()
		if err != nil {
			s.logger.Error().Err(err).Msg("Failed to start auth server")
			s.err <- err
		}
	}()
	defer func() {
		err := s.app.Shutdown()
		if err != nil {
			s.logger.Error().Err(err).Msg("Failed to shutdown auth server")
		}
		s.wg.Wait()
	}()
	select {
	case err := <-s.err:
		return nil, err
	case auth := <-s.auth:
		return auth, nil
	}
}

func (s *server) init() {
	s.app.Use(cors.New())
	s.app.Get(CallbackPath, func(ctx *fiber.Ctx) error {
		username := ctx.Query("username")
		if username == "" {
			return ctx.Status(fiber.StatusBadRequest).SendString("username is required")
		}
		accessToken := ctx.Query("access_token")
		if accessToken == "" {
			return ctx.Status(fiber.StatusBadRequest).SendString("access_token is required")
		}
		tokenType := ctx.Query("token_type")
		if tokenType == "" {
			return ctx.Status(fiber.StatusBadRequest).SendString("token_type is required")
		}
		refreshToken := ctx.Query("refresh_token")
		if refreshToken == "" {
			return ctx.Status(fiber.StatusBadRequest).SendString("refresh_token is required")
		}
		expiresIn := ctx.Query("expires_in")
		if expiresIn == "" {
			return ctx.Status(fiber.StatusBadRequest).SendString("expires_in is required")
		}

		parsedExpiresIn, err := strconv.Atoi(expiresIn)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString("expires_in must be an integer")
		}

		s.auth <- &Auth{
			Username:     username,
			AccessToken:  accessToken,
			TokenType:    tokenType,
			ExpiresIn:    parsedExpiresIn,
			RefreshToken: refreshToken,
		}

		return ctx.SendString("Authentication successful! You can close this window now.")
	})
}

func (s *server) start() error {
	var err error
	s.listener, err = net.Listen("tcp", Listen)
	if err != nil {
		return err
	}
	return s.app.Listener(s.listener)
}
