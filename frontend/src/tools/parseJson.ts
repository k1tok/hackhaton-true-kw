export async function parseJson(data) {
	const file = !!data.file?.[0] ? data.file?.[0] : null;
	let jsonData = null;

	if (file && file.type === "application/json") {
		const text = await file.text();
		try {
			jsonData = JSON.parse(text);
			console.log(jsonData);
		} catch (e) {
			console.error("Неверный JSON файл");
		}
	}

	return jsonData;
}
