import { Router, Route } from '@solidjs/router';
import UsersPage from './pages/UsersPage';
import HomePage from './pages/HomePage';
import AuthGuard from './components/AuthGuard';
import LoginPage from './pages/LoginPage';
import Layout from './pages/Layout';
import CreateUserPage from './pages/CreateUserPage';

export default function App() {
	return (
		<Router>
			<Route path="/" component={Layout}>
				<Route path="/" component={AuthGuard}>
					<Route path="/" component={HomePage} />
					<Route path="/users" component={UsersPage} />
				</Route>
			</Route>
			<Route path="/login" component={LoginPage} />
			<Route path="/createuser" component={CreateUserPage} />
		</Router>
	);
}
