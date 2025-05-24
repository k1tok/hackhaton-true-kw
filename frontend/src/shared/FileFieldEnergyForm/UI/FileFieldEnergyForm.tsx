import type { FieldValues } from "react-hook-form";
import type { IFileFieldEnergyForm } from "../config/fileFieldEnergyFormConfig";
import styles from "../style/FileFieldEnergyForm.module.css";

const FileFieldEnergyForm = <T extends FieldValues>({
	register,
	field,
	title,
}: IFileFieldEnergyForm<T>) => {
	return (
		<div className={styles.mainBlock}>
			<label className={styles.title}>{title}</label>
			<input {...register(field)} type="file" />
		</div>
	);
};

export default FileFieldEnergyForm;
