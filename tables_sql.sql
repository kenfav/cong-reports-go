CREATE TABLE publishers (
id SERIAL PRIMARY KEY,
surname varchar(50) not null,
name varchar(50) not null,
birthdate int not null,
baptismdate int null,
othersheep boolean DEFAULT true,
elder boolean DEFAULT false,
ministerialservant boolean DEFAULT false,
regularpionner boolean DEFAULT false,

)

CREATE TABLE reports (
  id SERIAL PRIMARY KEY,
  publishers_id INT NOT NULL,
  date_year int NOT NULL,
  date_month int NOT NULL,
  publications int null,
  videos int null,
  hours int not null,
  return_visits int null,
  bible_studies int null,
  notes varchar(200) null,
  CONSTRAINT fk_publishers FOREIGN KEY(publishers_id) REFERENCES publishers(id)
)
