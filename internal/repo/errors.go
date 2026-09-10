package repo

import "errors"

// ErrNotFound reports a missing row. GORM's own error is translated into it
// here so no other layer needs to import GORM.
var ErrNotFound = errors.New("data tidak ditemukan")
