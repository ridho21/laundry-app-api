package utils

const (

	//_________________________________________________ Master Data Customer
	INSERT_CUSTOMER = `INSERT INTO mst_customers(id,name,address,phone_number,email) VALUES ($1, $2, $3, $4, $5)`
	SELECT_CUSTOMER = `SELECT * FROM mst_customers`
	SELECT_CUSTOMER_PAGING = `SELECT * FROM mst_customers LIMIT $1 OFFSET $2`
	SELECT_CUSTOMER_ID = `SELECT * FROM mst_customers WHERE ID=$1`
	DELETE_CUSTOMER = `DELETE FROM mst_customers WHERE ID=$1`
	UPDATE_CUSTOMER = `UPDATE mst_customers SET name=$1, address=$2, phone_number=$3, email=$4 WHERE id=$5`
	SELECT_COUNT_CUSTOMER = `SELECT COUNT(*) FROM mst_customers`

	//_________________________________________________ Master Data Services
	SELECT_SERVICE = `SELECT * FROM mst_services`
	INSERT_SERVICE = `INSERT INTO mst_services(id,description,price) VALUES ($1, $2, $3)`
	SELECT_SERVICE_ID = `SELECT * FROM mst_services WHERE ID=$1`
	DELETE_SERVICE = `DELETE FROM mst_services WHERE ID=$1`
	UPDATE_SERVICE = `UPDATE mst_services SET description=$1, price=$2 WHERE id=$3`

	//_________________________________________________ Transaction
	INSERT_TRANSACTION = `INSERT INTO transactions (id,customer_id,pickup_date,status) VALUES ($1,$2,$3,$4) RETURNING id,customer_id,order_date,pickup_date,status`
	INSERT_TRX_DETAILS = `INSERT INTO transaction_details (id,transaction_id,service_id,qty) VALUES ($1,$2,$3,$4) RETURNING id,transaction_id,service_id,qty`

	//_________________________________________________ users
	INSERT_USER = `INSERT INTO users (id,full_name,email,username,password,role) VALUES ($1,$2,$3,$4,$5,$6) 
	RETURNING id,full_name,email,username,password,role`
	SELECT_USER_USERNAME = `SELECT * from users WHERE username=$1`
	SELECT_USER_ID = `SELECT * from users WHERE id=$1`
)

