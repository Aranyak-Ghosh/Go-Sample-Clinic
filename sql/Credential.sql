-- ActionName: EnsureCredentialCreated
CREATE TABLE IF NOT EXISTS `credential` (
		`id` VARCHAR(36) NOT NULL,
		`email` VARCHAR(255) NOT NULL,
		`hash_password` VARCHAR(4095) NOT NULL,
		PRIMARY KEY (`id`),
		UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE, 
		UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE);

-- ActionName: AddUserCredentials
INSERT INTO `techunicorn_clinic`.`credential`
(`id`,
`email`,
`hash_password`)
VALUES
(<{id:}>,
<{email:}>,
<{hash_password:}>);

-- ActionName: FetchUserCredentialsById
SELECT `id`, `email`, `hash_password`
FROM `credential`
WHERE `id` = <{id: }>
LIMIT 1

-- ActionName: FetchUserCredentialsByEmail
SELECT `id`, `email`, `hash_password`
FROM `credential`
WHERE `email` = <{email: }>
LIMIT 1