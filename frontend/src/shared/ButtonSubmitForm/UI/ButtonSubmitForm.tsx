import styles from "../style/ButtonSubmitForm.module.css";

interface IButtonSubmitForm {
	title: string;
}

const ButtonSubmitForm = ({ title }: IButtonSubmitForm) => {
	return (
		<div className={styles.mainBlock}>
			<input className={styles.btn} type="submit" value={title} />
		</div>
	);
};

export default ButtonSubmitForm;
