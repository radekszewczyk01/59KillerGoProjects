# **Projekt: Go & MySQL - System do zarządzania książkami**

### Opis projektu
Ten projekt w języku Go umożliwia zarządzanie książkami w bazie danych MySQL. Aplikacja łączy się z bazą danych, pozwala na tworzenie, usuwanie i przeglądanie książek poprzez wysyłanie zapytań HTTP do lokalnego serwera. Projekt obsługuje CRUD (Create, Read, Update, Delete) dla rekordów w tabeli `books`.

### Technologie:
- **Go** - język programowania wykorzystywany do tworzenia aplikacji.
- **MySQL** - baza danych, w której przechowywane są informacje o książkach.
- **GORM** - ORM (Object Relational Mapping) dla Go, ułatwiający interakcję z bazą danych.
- **HTTP** - serwer, który umożliwia komunikację z aplikacją za pomocą zapytań HTTP.
  
### Funkcje:
- **POST**: Dodawanie nowych książek do bazy danych.
- **DELETE**: Usuwanie książek z bazy danych po identyfikatorze.

### Jak to działa?

1. **Połączenie z bazą danych**: 
   Projekt łączy się z bazą MySQL przy pomocy GORM-a. Dane o książkach są przechowywane w tabeli `books`, która zawiera następujące kolumny:
   - `id` (AUTO_INCREMENT, Primary Key) – unikalny identyfikator książki.
   - `name` – nazwa książki.
   - `author` – autor książki.
   - `publication` – wydawnictwo książki.
   - `created_at`, `updated_at` – daty utworzenia i modyfikacji rekordu.
   - `deleted_at` – data usunięcia (jeśli książka została usunięta, używamy tzw. twardego usuwania).

2. **Wysyłanie żądań do aplikacji**: 
   Aplikacja nasłuchuje na porcie `9010` i obsługuje zapytania HTTP. Przykładowo, możesz dodać książki za pomocą `POST` i usunąć je za pomocą `DELETE`.

---

### Instalacja

1. **Zainstaluj MySQL**:
   Upewnij się, że masz zainstalowaną bazę danych MySQL oraz że masz dostęp do niej na porcie 3306.

2. **Zainstaluj Go**:
   Pobierz i zainstaluj Go: https://golang.org/dl/

3. **Stwórz bazę danych**:
   Utwórz bazę danych o nazwie `simplerest` w MySQL:
   ```sql
   CREATE DATABASE simplerest;
   ```

4. **Zainstaluj zależności Go**:
   Zainstaluj wymagane pakiety przy użyciu polecenia:
   ```bash
   go mod tidy
   ```

5. **Uruchom aplikację**:
   Aby uruchomić aplikację Go, użyj poniższego polecenia:
   ```bash
   go run main.go
   ```

### Przykład użycia API

Po uruchomieniu aplikacji, API będzie dostępne na `http://localhost:9010/`.

1. **Dodanie książki**:
   Aby dodać książkę do bazy, wyślij żądanie `POST` na endpoint `http://localhost:9010/book/` z danymi książki:

   **Przykładowe zapytanie w JavaScript (fetch):**

   ```javascript
   fetch("http://localhost:9010/book/", {
       method: "POST",
       headers: {
           "Content-Type": "application/json"
       },
       body: JSON.stringify({
           "Name": "Lord of The Rings",
           "Author": "J.R.R. Tolkien",
           "Publication": "XYZ"
       })
   })
   .then(response => response.json())
   .then(data => console.log(data))
   .catch(error => console.error("Błąd:", error));
   ```

   **Inny przykład książki:**

   ```javascript
   fetch("http://localhost:9010/book/", {
       method: "POST",
       headers: {
           "Content-Type": "application/json"
       },
       body: JSON.stringify({
           "Name": "Wiedźmin",
           "Author": "Andrzej Sapkowski",
           "Publication": "SuperNOWA"
       })
   })
   .then(response => response.json())
   .then(data => console.log(data))
   .catch(error => console.error("Błąd:", error));
   ```

   **Oczekiwana odpowiedź JSON:**

   ```json
   {
       "id": 10,
       "name": "Lord of The Rings",
       "author": "J.R.R. Tolkien",
       "publication": "XYZ",
       "created_at": "2025-03-31T22:16:27Z",
       "updated_at": "2025-03-31T22:16:27Z",
       "deleted_at": null
   }
   ```

2. **Usuwanie książki**:
   Aby usunąć książkę z bazy, wyślij zapytanie `DELETE` na endpoint `http://localhost:9010/book/{id}`.

   **Przykładowe zapytanie w JavaScript:**

   ```javascript
   fetch('http://localhost:9010/book/1', {
     method: 'DELETE',
   })
   .then(response => {
     if (response.ok) {
       console.log('Książka została usunięta');
     } else {
       console.log('Błąd przy usuwaniu książki');
     }
   })
   .catch(error => console.error('Błąd: ', error));
   ```

   **Oczekiwana odpowiedź:**

   ```json
   {
       "message": "Książka została usunięta"
   }
   ```

---

### Przykład stanu tabeli po dodaniu książek:

Po wykonaniu zapytania `SELECT * FROM books;`, możesz zobaczyć w tabeli dane o książkach, takie jak:

```sql
+----+-------------------+-------------------+-------------+---------------------+---------------------+------------+
| id | name              | author            | publication | created_at          | updated_at          | deleted_at |
+----+-------------------+-------------------+-------------+---------------------+---------------------+------------+
| 10 | Lord of The Rings | J.R.R. Tolkien    | XYZ         | 2025-03-31 22:16:27 | 2025-03-31 22:16:27 | NULL       |
| 11 | Wiedźmin          | Andrzej Sapkowski | SuperNOWA   | 2025-03-31 22:16:52 | 2025-03-31 22:16:52 | NULL       |
+----+-------------------+-------------------+-------------+---------------------+---------------------+------------+
```

---

### Uwagi:
- Projekt korzysta z **soft delete** w GORM, co oznacza, że rekordy są "usuwane" tylko przez ustawienie pola `deleted_at`, zamiast usuwania ich fizycznie z bazy.
- Każda książka ma pole `created_at` oraz `updated_at`, które są automatycznie ustawiane przez GORM.

---

### Licencja

Projekt jest udostępniony na licencji MIT.

