# Embed binary or text files into your executable.

Sometimes one wants to avoid dealin with direct filesystem communication, but have error proof access to a named set of bytes. One of the ways is to embedd the binary stream into the executable.

How to use this tool
{{{ 
//go:generate resource <yourproject/res/resource.go> [input_file1] [input_file2] [etc...]
}}}
then inside of your code just
{{{
res.Resources["input_file1"] to access the byte array with the file contents
}}}
