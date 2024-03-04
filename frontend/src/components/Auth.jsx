export default function AuthForm({ title='Auth', children }) {
	document.title = title
	return (
		<div className="px-6 py-6 max-w-sm mx-auto bg-zinc-800 text-zinc-400 rounded">
			<div>
				<h1 className="text-center">
					{title}
				</h1>
				{children}
			</div>
		</div>
	);
}
