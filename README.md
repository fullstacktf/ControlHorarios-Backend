## Tecnologías utilizadas

---

* Go
* Nginx
* MariaDB


## Endpoints

---

### Company

---

| URI                        | Request |
|----------------------------|---------|
| /api/company               | POST    |
| /api/company/login         | POST    |
| /api/company/settings/{id} | POST    |
| /api/company/settings/{id} | GET     |
| /api/company/calendar      | POST    |
| /api/company/calendar      | GET     |

### Employee

---

| URI                         | Request |
|-----------------------------|---------|
| /api/employee/login         | POST    |
| /api/employee/{id}          | GET     |
| /api/employee/calendar/{id} | GET     |
| /api/employee/summary/{id}  | GET     |
| /api/employee/settings/{id} | GET     |
| /api/employee/settings/{id} | POST    |


## Database Schema

---

![Database Schema](/docs/images/bbdd.PNG ";)")