package prototype_file

type Inode interface {
	print(string)
	clone() Inode
}
