import type { IFormInput } from "../entities/EnergyForm/config/energyFormConfig";

export function fetchData(url: string, data: IFormInput[]) {
	return fetch(url, {
		method: "POST",
		body: JSON.stringify(data),
		headers: {
			"Content-Type": "application/json",
		},
	})
		.then((response) => {
			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}

			return response.json();
		})
		.catch((e) => {
			console.error("Fetch error:", {
				url,
				error: e.message,
				payload: data,
			});
			throw e;
		});
}
