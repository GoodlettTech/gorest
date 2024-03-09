import { jwt } from '../signals/jwt';

export default function HomePage() {
	return (
		<div>
			<h1>Home Page {jwt()}</h1>
		</div>
	);
}
