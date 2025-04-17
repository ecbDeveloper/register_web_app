export function setItemWithExpiry(id: string) {
	const now = new Date();
	const item = {
		id,
		expiry: now.getTime() + 1000 * 60 * 60,
	};
	localStorage.setItem("userId", JSON.stringify(item));
}

export function getItemWithExpiry(): string | null {
	const userId = localStorage.getItem("userId");
	if (!userId) return null;

	try {
		const userIdJSON = JSON.parse(userId);
		const now = new Date();
		if (now.getTime() > userIdJSON.expiry) {
			localStorage.removeItem("userId");
			return null;
		}

		return userIdJSON.id;
	} catch {
		localStorage.removeItem("userId");
		return null;
	}
}