package models

import "time"

type User struct {
	ID        int       
	Name      string    
	Email     string    
	CreatedAt time.Time 
}

func (u *User) Validate() error {
	return nil
}
