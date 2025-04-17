export function TruncateText(str: string) {
	if (str.length > 24) {
		return str.slice(0, 24) + '...';
	}
	return str;
}
