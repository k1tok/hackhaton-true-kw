export interface IFormInput {
	accountId: number;
	address: string;
	buildingType: string;
	roomsCount?: number;
	residentsCount?: number;
	totalArea?: number;
	consumption: {
		[key: string]: number;
	};
	file?: FileList;
}

export const isEmptyObject = (obj: any) => {
	return Object.entries(obj)
		.filter(([key]) => !["file", "anotherFieldToIgnore"].includes(key))
		.every(
			([_, value]) => value === "" || value === null || value === undefined
		);
};
