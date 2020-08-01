首先启动Java虚拟机，然后加载主类，最后调用主类的main()方法。
```java
public class HelloWorld() {

    public static void main(String []args) {
        System.out.println("Hello, Tomonori!");
    }
}
```
加载HelloWorld类之前，首先要加载他的超类，也就是java.lang.Object。在调用main函数之前，因为虚拟机需要准备好餐参数数组，所以需要加载java.lang.Object和java.lang.String[]类。
把字符串打印到控制台还需要加载java.lang.System类。

Oracle的Java虚拟机实现根据路劲(class path)来搜索类。按照搜索的先后顺序，类路劲可以分为以下3个部分:
- 启动类路径(bootstrap classpath)：默认对应jre/lib目录，Java标准库(大部分在rt.jar里)位于该路径。可以通过-Xbootclasspath选项修改路径，不需要这样做
- 扩展类路径(extension classpath)：默认对应jre/lib/ext目录，使用Java扩展机制的类位于这个路径。
- 用户类路径(user classpath)：我们自己实现的类，以及第三方类库则位于用户类路径。默认值是：. 可以设置CLASSPATH环境变量修改用户类路径，但这不灵活，更好的是给java命令传递-classpath(或简写-cp)，优先级更高，覆盖环境变量。-classpath/-cp选项既可以指定目录，也可以指定JAR文件或ZIP文件。还可以同时指定多个目录或文件，用分隔符分开即可，win中式分号，unix是冒号

