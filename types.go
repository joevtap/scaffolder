package scaffolder

// Project struct represents the project config file.
type Project struct {
	Root []Block `toml:"root" validate:"required,dive"`
}

// Block struct represents a directory or file.
//
//  - Name is the name of the directory or file.
//  - Type is the type of the block. It can be either "dir" or "file".
//  - Template is the path to the template file to be used when rendering the file.
//   	The path is relative to the App Root, and is only required if the block type is "file".
//  - Children is a slice of Block structs that represent the children of the directory.
//	 	Children is only required if the block type is "dir".
type Block struct {
	Name     string  `toml:"name" validate:"required,min=3"`
	Type     string  `toml:"type" validate:"required,oneof=dir file"`
	Template string  `toml:"template,omitempty" validate:"omitempty,required_if=Type file"`
	Children []Block `toml:"children,omitempty" validate:"omitempty,dive"`
}
