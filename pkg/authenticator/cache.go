package authenticator

import (
	"crypto/sha256"
	"encoding/hex"
	"sync"
	"time"

	"k8s.io/apiserver/pkg/authentication/user"
)

// Cache est un cache thread-safe pour les utilisateurs authentifiés
type Cache struct {
	mu      sync.RWMutex
	entries map[string]*cacheEntry
	ttl     time.Duration
}

type cacheEntry struct {
	user      *user.DefaultInfo
	expiresAt time.Time
}

// NewCache crée un nouveau cache avec le TTL spécifié
func NewCache(ttl time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]*cacheEntry),
		ttl:     ttl,
	}

	// Lancer le nettoyage automatique des entrées expirées
	go c.cleanupExpired()

	return c
}

// Get récupère un utilisateur du cache
func (c *Cache) Get(walletAddress string) (*user.DefaultInfo, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	key := c.getCacheKey(walletAddress)
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	if time.Now().After(entry.expiresAt) {
		return nil, false
	}

	return entry.user, true
}

// Set ajoute un utilisateur au cache
func (c *Cache) Set(walletAddress string, userInfo *user.DefaultInfo) {
	c.mu.Lock()
	defer c.mu.Unlock()

	key := c.getCacheKey(walletAddress)
	c.entries[key] = &cacheEntry{
		user:      userInfo,
		expiresAt: time.Now().Add(c.ttl),
	}
}

// Delete supprime un utilisateur du cache
func (c *Cache) Delete(walletAddress string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	key := c.getCacheKey(walletAddress)
	delete(c.entries, key)
}

// Clear vide complètement le cache
func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries = make(map[string]*cacheEntry)
}

// Size retourne le nombre d'entrées dans le cache
func (c *Cache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.entries)
}

// getCacheKey génère une clé de cache à partir d'une adresse wallet
func (c *Cache) getCacheKey(wallet string) string {
	hash := sha256.Sum256([]byte(wallet))
	return hex.EncodeToString(hash[:])
}

// cleanupExpired supprime périodiquement les entrées expirées
func (c *Cache) cleanupExpired() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, entry := range c.entries {
			if now.After(entry.expiresAt) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
