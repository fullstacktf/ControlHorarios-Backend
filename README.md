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
| /api/companies               | POST    |
| /api/companies/login         | POST    |
| /api/companies/{id}/holidays | POST    |
| /api/companies/{id}/holidays | GET     |



### Employee

---

| URI                         | Request |
|-----------------------------|---------|
| /api/employees/login         | POST    |
| /api/employees/{id}          | GET     |
| /api/employees/{id}/summary  | GET     |
| /api/employees               | POST    |
| /api/employees/{id}/password | PUT     |
| /api/employees/{id}/checkin  | POST    |
| /api/employees/{id}/checkout | POST    |


## Database Schema

---

![Database Schema](/docs/images/bbdd.PNG ";)")
