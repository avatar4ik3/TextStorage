export default class Group {
	private _Id: Number
	public get Id(): Number {
		return this._Id
	}
	public set Id(v: Number) {
		this._Id = v
	}

	private _Description: string
	public get Description(): string {
		return this._Description
	}
	public set Description(v: string) {
		this._Description = v
	}

	constructor(id: Number, description: string) {
		this._Id = id
		this._Description = description
	}
}
