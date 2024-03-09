import { useNavigate } from '@solidjs/router';
import { createEffect } from 'solid-js';

import { jwt } from '../signals/jwt';

export default function AuthGuard(props) {
	const navigate = useNavigate();

	createEffect(() => {
		if (!jwt()) {
			navigate('/login', {
				replace: true,
			});
		}
	});

	return <>{props.children}</>;
}
