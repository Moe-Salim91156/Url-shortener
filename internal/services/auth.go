package services

import (
	"Url-shortener/internal/models"
	"Url-shortener/internal/store"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)
type 
