const defineUserData = async (value) => {
	const user = {
		firstName: value.first_name,
		lastName: value.last_name,
		email: value.email,
		token: value.token,
		cpf: value.cpf,
		craetedAt: value.created_at,
		updatedAt: value.updated_at,
	}

	localStorage.setItem("user", JSON.stringify(user))

};

export default defineUserData;
