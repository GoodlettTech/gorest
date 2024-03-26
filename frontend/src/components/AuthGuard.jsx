import { useNavigate, useLocation } from '@solidjs/router';
import { createEffect } from 'solid-js';

import { jwt } from '../signals/jwt';

export default function AuthGuard(props) {
	const navigate = useNavigate();
	
	createEffect(() => {
		if (!jwt()) {
			const location = useLocation();
			navigate(`/login?redirect=${location.pathname}${location.search}`, {
				replace: true,
			});
		}
	});

	return <>{props.children}</>;
}
