import Navbar from '../components/Navbar';
export default function Layout(props) {
	return (
		<div class="container-fluid px-0">
			<Navbar></Navbar>
			{props.children}
		</div>
	);
}
