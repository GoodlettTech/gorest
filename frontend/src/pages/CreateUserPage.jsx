import AuthForm from "../components/Auth";
import CreateUserForm from "../components/CreateUserForm";

export default function CreateUserPage() {
	return (
		<AuthForm title="Create User">
            <CreateUserForm/>
        </AuthForm>
	);
}