import "./App.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Home from "./pages/Home";
import LoginPage from "./pages/LoginPage";
import CreateUserPage from "./pages/CreateUserPage";

const router = createBrowserRouter([
	{
		path: "/",
		element: <Home />,
	},
	{
		path: "auth",
		children: [
			{
				path:"login",
				element: <LoginPage/>
			},
			{
				path:"createuser",
				element: <CreateUserPage />
			}
		]
	}
]);

function App() {
	return (
		<div className="container-lg">
			<RouterProvider router={router} />
		</div>
	);
}

export default App;
