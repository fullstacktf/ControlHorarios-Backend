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


### Employee

---

| URI                         | Request |
|-----------------------------|---------|
| /api/employee/login         | POST    |
| /api/employee/{id}          | GET     |
| /api/employee/summary/{id}  | GET     |
| /api/employee               | POST    |
| /api/employee/password/{id} | PUT     |
| /api/employee/checkin/{id}  | POST    |
| /api/employee/checkout/{id} | POST    |

### Holidays

| URI                        | Request |
|----------------------------|---------|
| /api/company/holidays/{id} | POST    |
| /api/company/holidays/{id} | GET     |

## Database Schema

---

![Database Schema](https://i.imgur.com/wqSHswx.png)
