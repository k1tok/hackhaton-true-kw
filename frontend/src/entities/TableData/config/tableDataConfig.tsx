export const headersTable = [
	{ val: "Номер обращения" },
	{ val: "Адрес" },
	{ val: "Количество комнат" },
	{ val: "Количество жильцов" },
	{ val: "Площадь помещения" },
	{ val: "Тип строения" },
	{ val: "Потребление энергии" },
];

export const getMonthName = (monthNumber) => {
	const months = [
		"Январь",
		"Февраль",
		"Март",
		"Апрель",
		"Май",
		"Июнь",
		"Июль",
		"Август",
		"Сентябрь",
		"Октябрь",
		"Ноябрь",
		"Декабрь",
	];
	return months[monthNumber - 1] || monthNumber;
};

export const renderConsumption = (consumption) => {
	if (!consumption || Object.keys(consumption).length === 0) {
		return <div style={{ color: "#999" }}>Нет данных</div>;
	}

	return (
		<div
			style={{
				display: "grid",
				gridTemplateColumns: "repeat(3, 1fr)",
				gap: "8px",
			}}
		>
			{Object.entries(consumption)
				.sort(([a], [b]) => parseInt(a) - parseInt(b))
				.map(([month, value]) => (
					<div key={month} style={{ whiteSpace: "nowrap" }}>
						<strong>{getMonthName(month)}:</strong> {value}
					</div>
				))}
		</div>
	);
};
