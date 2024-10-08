import { toast } from 'svelte-sonner';

export function showError(errCode: string) {
	let error = errCode;
	let desc = '';

	switch (errCode) {
		case 'no-user-exists':
			error = 'Username Not Found';
			desc = 'This user has not been registered';
			break;
		case 'wrong-password':
			error = 'Wrong Password';
			desc = 'The password you entered is incorrect';
			break;
		case 'passwords-dont-match':
			error = "Passwords Don't Match";
			desc = 'The passwords you entered do not match each other';
			break;
	}

	toast.error(error, {
		description: desc,
		action: {
			label: 'Close',
			onClick: () => {}
		}
	});
}

export function showSuccess(title: string, desc: string) {
	toast.success(title, {
		description: desc,
		class: 'bg-red-900',
		action: {
			label: 'Close',
			onClick: () => {}
		}
	});
}
