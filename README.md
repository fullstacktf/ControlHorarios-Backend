## Tecnologías utilizadas

---

* Go
* Nginx
* MariaDB



## Endpoints

---

### Company

---

| URI                        | Request | Body |
|----------------------------|---------|------|
| /api/companies               | POST    | username, email, password, company_name, location |
| /api/companies/login         | POST    | email, password |
| /api/companies/{id}/holidays | POST    | holiday_title, holiday_date |
| /api/companies/{id}/holidays | GET     |  |



### Employee

---

| URI                         | Request | Body |
|-----------------------------|---------|------|
| /api/employees               | POST    | username, email, password, first_name, last_name, company_id |
| /api/employees/login         | POST    | email, password |
| /api/employees/{id}          | GET     | |
| /api/employees/{id}/summary  | GET     | |
| /api/employees/{id}/password | PUT     | password |
| /api/employees/{id}/checkin  | POST    | description |
| /api/employees/{id}/checkout | POST    | |


## Database Schema

---

![Database Schema](/docs/images/bbdd.png ";)")
