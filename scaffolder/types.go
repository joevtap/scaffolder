package scaffolder

type Project struct {
	Root []Block `toml:"root" validate:"required,dive"`
}

type Block struct {
	Name     string  `toml:"name" validate:"required,min=3"`
	Type     string  `toml:"type" validate:"required,oneof=dir file"`
	Template string  `toml:"template,omitempty" validate:"omitempty,required_if=Type file,file"`
	Children []Block `toml:"children,omitempty" validate:"omitempty,dive"`
}
