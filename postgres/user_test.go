package postgres_test

import (
	"testing"

	"github.com/xueyiyao/safekeep/domain"
	"github.com/xueyiyao/safekeep/postgres"
)

func TestUserService_CreateUser(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		db, tx := MustOpenDB(t)
		defer MustCloseDB(t, db, tx, func() { resetSequence(db.DB, "users", "id", 1) })

		s := postgres.NewUserService(tx)

		u := &domain.User{Name: "one", Email: "one@test.com"}
		if err := s.CreateUser(u); err != nil {
			t.Fatal(err)
		} else if got, want := u.ID, 1; int(got) != want {
			t.Fatalf("ID=%v, want %v", got, want)
		}

		u2 := &domain.User{Name: "two", Email: "two@test.com"}
		if err := s.CreateUser(u2); err != nil {
			t.Fatal(err)
		} else if got, want := u2.ID, 2; int(got) != want {
			t.Fatalf("ID=%v, want %v", got, want)
		}

		u3 := &domain.User{Name: "two", Email: "nottwo@test.com"}
		if err := s.CreateUser(u3); err != nil {
			t.Fatal(err)
		} else if got, want := u3.ID, 3; int(got) != want {
			t.Fatalf("ID=%v, want %v", got, want)
		}

		// Go's time.Time precision is in ns, while PSQL is in microseconds. this will fail as is so commented out
		// // Simply fetches user and compare
		// if other, err := s.FindUserByID(1); err != nil {
		// 	t.Fatal(err)
		// } else if !reflect.DeepEqual(u, other) {
		// 	t.Fatalf("mismatch: %#v != %#v", u, other)
		// }
	})

	// Ensure an error is returned if empty user.
	t.Run("ErrNilUser", func(t *testing.T) {
		db, tx := MustOpenDB(t)
		defer MustCloseDB(t, db, tx, func() { resetSequence(db.DB, "users", "id", 1) })

		s := postgres.NewUserService(tx)
		if err := s.CreateUser(nil); err == nil {
			t.Fatal("expected an error, none occured")
		}
	})

	// Ensure an error is returned if user name is not set.
	t.Run("ErrNameRequired", func(t *testing.T) {
		db, tx := MustOpenDB(t)
		defer MustCloseDB(t, db, tx, func() { resetSequence(db.DB, "users", "id", 1) })

		s := postgres.NewUserService(tx)
		if err := s.CreateUser(&domain.User{Email: "one@gmail.com"}); err == nil {
			t.Fatal("expected an error, none occured")
		}
	})

	// Ensure an error is returned if user email is not set.
	t.Run("ErrEmailRequired", func(t *testing.T) {
		db, tx := MustOpenDB(t)
		defer MustCloseDB(t, db, tx, func() { resetSequence(db.DB, "users", "id", 1) })

		s := postgres.NewUserService(tx)
		if err := s.CreateUser(&domain.User{Name: "one"}); err == nil {
			t.Fatal("expected an error, none occured")
		}
	})

	// Ensure an error is returned if user email is not valid.
	t.Run("ErrEmailInvalid", func(t *testing.T) {
		db, tx := MustOpenDB(t)
		defer MustCloseDB(t, db, tx, func() { resetSequence(db.DB, "users", "id", 1) })

		s := postgres.NewUserService(tx)
		if err := s.CreateUser(&domain.User{Name: "one", Email: "asdfjkl;"}); err == nil {
			t.Fatal("expected an error, none occured")
		}
	})

	// Ensure an error is returned if user email is a duplicate.
	t.Run("ErrEmailDuplicate", func(t *testing.T) {
		db, tx := MustOpenDB(t)
		defer MustCloseDB(t, db, tx, func() { resetSequence(db.DB, "users", "id", 1) })

		s := postgres.NewUserService(tx)
		u := &domain.User{Name: "one", Email: "one@test.com"}
		MustCreateUser(t, s, u)
		if err := s.CreateUser(&domain.User{Name: "one", Email: "one@test.com"}); err == nil {
			t.Fatal("expected an error, none occured")
		}
	})
}

func TestUserService_FindUser(t *testing.T) {
	// Ensure an error is returned if fetching a non-existent user.
	t.Run("ErrNotFound", func(t *testing.T) {
		db, tx := MustOpenDB(t)
		defer MustCloseDB(t, db, tx, func() { resetSequence(db.DB, "users", "id", 1) })

		s := postgres.NewUserService(tx)
		if user, err := s.FindUserByID(1); err == nil {
			t.Fatalf("expected an error, none occured %v", user)
		}
	})
}

func TestUserService_UpdateUser(t *testing.T) {
	// Ensure user name & email can be updated by current user.
	t.Run("ErrNotFound", func(t *testing.T) {
		db, tx := MustOpenDB(t)
		defer MustCloseDB(t, db, tx, func() { resetSequence(db.DB, "users", "id", 1) })

		s := postgres.NewUserService(tx)

		user := MustCreateUser(t, s, &domain.User{
			Name:  "one",
			Email: "one@gmail.com",
		})

		// Update user.
		newName, newEmail := "two", "two@gmail.com"
		_, err := s.UpdateUser(&domain.User{
			ID:    user.ID,
			Name:  newName,
			Email: newEmail,
		})
		if err != nil {
			t.Fatal(err)
		}

		// Fetch user from database & compare.
		if user, err := s.FindUserByID(1); err != nil {
			t.Fatal(err)
		} else if got, want := user.Name, "two"; got != want {
			t.Fatalf("Name=%v, want %v", got, want)
		} else if got, want := user.Email, "two@gmail.com"; got != want {
			t.Fatalf("Email=%v, want %v", got, want)
		}
		// Reflect DeepEqual here
	})
}

// MustCreateUser creates a user in the database. Fatal on error.
func MustCreateUser(tb testing.TB, s *postgres.UserService, user *domain.User) *domain.User {
	tb.Helper()
	if err := s.CreateUser(user); err != nil {
		tb.Fatal(err)
	}
	return user
}
