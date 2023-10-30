CREATE TABLE brand  (
  id int(11) NOT NULL AUTO_INCREMENT,
  title varchar(100) NOT NULL,
  createdAt datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  updatedAt datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (id)
) ENGINE = InnoDB;