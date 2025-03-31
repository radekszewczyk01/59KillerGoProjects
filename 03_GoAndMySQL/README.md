# **Project: Go & MySQL - Book Management System**

### Project Description
This project, written in Go, allows managing books in a MySQL database. The application connects to the database and enables creating, deleting, and viewing books by sending HTTP requests to a local server. The project supports full CRUD (Create, Read, Update, Delete) operations for records in the `books` table.

### Technologies:
- **Go** - The programming language used to develop the application.
- **MySQL** - The database used to store book information.
- **GORM** - An ORM (Object Relational Mapping) for Go, simplifying database interactions.
- **HTTP** - A server that enables communication with the application through HTTP requests.
  
### Features:
- **POST**: Adds new books to the database.
- **DELETE**: Removes books from the database by ID.

### How It Works?

1. **Database Connection**:
   The project connects to a MySQL database using GORM. Book data is stored in the `books` table, which includes the following columns:
   - `id` (AUTO_INCREMENT, Primary Key) – Unique book identifier.
   - `name` – Name of the book.
   - `author` – Author of the book.
   - `publication` – Publisher of the book.
   - `created_at`, `updated_at` – Timestamps for record creation and modification.
   - `deleted_at` – Timestamp for deletion (soft delete is used).

2. **Sending Requests to the Application**:
   The application listens on port `9010` and handles HTTP requests. For example, you can add books using `POST` and delete them using `DELETE`.

---

### Installation

1. **Install MySQL**:
   Ensure you have MySQL installed and running on port 3306.

2. **Install Go**:
   Download and install Go from: https://golang.org/dl/

3. **Create the Database**:
   Create a MySQL database named `simplerest`:
   ```sql
   CREATE DATABASE simplerest;
   ```

4. **Install Go Dependencies**:
   Install required packages using:
   ```bash
   go mod tidy
   ```

5. **Run the Application**:
   Start the Go application using:
   ```bash
   go run main.go
   ```

### API Usage Examples

Once the application is running, the API will be available at `http://localhost:9010/`.

1. **Adding a Book**:
   To add a book to the database, send a `POST` request to `http://localhost:9010/book/` with book details:

   **Example JavaScript (fetch) request:**

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
   .catch(error => console.error("Error:", error));
   ```

   **Another Example:**

   ```javascript
   fetch("http://localhost:9010/book/", {
       method: "POST",
       headers: {
           "Content-Type": "application/json"
       },
       body: JSON.stringify({
           "Name": "The Witcher",
           "Author": "Andrzej Sapkowski",
           "Publication": "SuperNOWA"
       })
   })
   .then(response => response.json())
   .then(data => console.log(data))
   .catch(error => console.error("Error:", error));
   ```

   **Expected JSON Response:**

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

2. **Deleting a Book**:
   To delete a book from the database, send a `DELETE` request to `http://localhost:9010/book/{id}`.

   **Example JavaScript request:**

   ```javascript
   fetch('http://localhost:9010/book/1', {
     method: 'DELETE',
   })
   .then(response => {
     if (response.ok) {
       console.log('Book deleted successfully');
     } else {
       console.log('Error deleting the book');
     }
   })
   .catch(error => console.error('Error:', error));
   ```

   **Expected Response:**

   ```json
   {
       "message": "Book has been deleted"
   }
   ```

---

### Example Table State After Adding Books:

After running `SELECT * FROM books;`, you should see the following records:

```sql
+----+-------------------+-------------------+-------------+---------------------+---------------------+------------+
| id | name              | author            | publication | created_at          | updated_at          | deleted_at |
+----+-------------------+-------------------+-------------+---------------------+---------------------+------------+
| 10 | Lord of The Rings | J.R.R. Tolkien    | XYZ         | 2025-03-31 22:16:27 | 2025-03-31 22:16:27 | NULL       |
| 11 | The Witcher       | Andrzej Sapkowski | SuperNOWA   | 2025-03-31 22:16:52 | 2025-03-31 22:16:52 | NULL       |
+----+-------------------+-------------------+-------------+---------------------+---------------------+------------+
```

---

### Notes:
- The project uses **soft delete** in GORM, meaning records are "deleted" by setting the `deleted_at` field instead of being physically removed from the database.
- Each book record includes `created_at` and `updated_at` fields, which are automatically managed by GORM.

---

### License

This project is released under the MIT License.

