这个文件夹实现了书中的xmlselect.go。

test.xml为xmlselect.go的输入，output.txt为输出，test.html也可作为xmlselect.go的输入，但是会报错，因为html与XML有些区别，html不要求每个标签都是成对的，一开一闭，xml遇到这种情况会报错，停止解析。

xmlselect的具体功能为对某个XML文档进行解析，识别满足特定条件的有意义的文字内容。

具体的原理为，xml包解析XML时，会识别不同的文法结构，并保存为不同的类型。比如`<Textview layout_gravity="center"/>`被识别为StartElement，`</Textview>`被识别成EndElement，有意义的文本内容识别成CharData，等等。

