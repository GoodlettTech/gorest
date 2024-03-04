import { useEffect } from 'react';

export default function Home() {

	useEffect(() => {
		document.title = "Home Page"
	}, [])

	return (
		<div className="container-xl">
			test
		</div>
	)
}