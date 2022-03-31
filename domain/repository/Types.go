package repository

type PreparedStatement string

const (
	//\/\/\/\/\/\/\/\Credential Table/\/\/\/\/\/\/\/\\
	EnsureCredentialCreated PreparedStatement = "CREATE TABLE IF NOT EXISTS `credential` (\n" +
		"`id` VARCHAR(36) NOT NULL,\n" +
		"`email` VARCHAR(255) NOT NULL,\n" +
		"`hash_password` VARCHAR(4095) NOT NULL,\n" +
		"PRIMARY KEY (`id`),\n" +
		"UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,\n" +
		"UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE);"

	AddUserCredentials PreparedStatement = "INSERT INTO `techunicorn_clinic`.`credential` \n" +
		"(`id`, \n" +
		"`email`, \n" +
		" `hash_password`) \n" +
		"VALUES \n" +
		"(:id, \n" +
		":email, \n" +
		":hash_password);"
	FetchUserCredentialsById PreparedStatement = "SELECT `id`, `email`, `hash_password` \n" +
		"FROM `credential` \n" +
		"WHERE `id` = :id \n" +
		"LIMIT 1"
	FetchUserCredentialsByEmail PreparedStatement = "SELECT `id`, `email`, `hash_password` \n" +
		"FROM `credential` \n" +
		"WHERE `email` = :email \n" +
		"LIMIT 1"
)
