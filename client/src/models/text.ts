export default class Text {
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

	private _Value: string
	public get Value(): string {
		return this._Value
	}
	public set Value(v: string) {
		this._Value = v
	}

	constructor(id: Number, description: string, value: string) {
		this._Id = id
		this._Value = value
		this._Description = description
	}
}
