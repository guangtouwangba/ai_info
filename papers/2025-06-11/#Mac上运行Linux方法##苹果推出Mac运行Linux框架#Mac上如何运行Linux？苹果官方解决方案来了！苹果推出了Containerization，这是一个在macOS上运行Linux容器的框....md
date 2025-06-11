# #Mac上运行Linux方法##苹果推出Mac运行Linux框架#Mac上如何运行Linux？苹果官方解决方案来了！苹果推出了Containerization，这是一个在macOS上运行Linux容器的框...

**URL**: https://weibo.com/6105753431/Pw3qJBYSv

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Mac%E4%B8%8A%E8%BF%90%E8%A1%8CLinux%E6%96%B9%E6%B3%95%23&amp;extparam=%23Mac%E4%B8%8A%E8%BF%90%E8%A1%8CLinux%E6%96%B9%E6%B3%95%23" data-hide=""><span class="surl-text">#Mac上运行Linux方法#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%8B%B9%E6%9E%9C%E6%8E%A8%E5%87%BAMac%E8%BF%90%E8%A1%8CLinux%E6%A1%86%E6%9E%B6%23&amp;extparam=%23%E8%8B%B9%E6%9E%9C%E6%8E%A8%E5%87%BAMac%E8%BF%90%E8%A1%8CLinux%E6%A1%86%E6%9E%B6%23" data-hide=""><span class="surl-text">#苹果推出Mac运行Linux框架#</span></a><br><br>Mac上如何运行Linux？苹果官方解决方案来了！<br><br>苹果推出了Containerization，这是一个在macOS上运行Linux容器的框架。该技术使用Virtualization.framework，并优化了Apple Silicon上的表现。<br><br>构建与运行要求：<br><br>- 必须使用Apple Silicon的Mac。<br>    <br>- 构建容器化框架需要macOS 15及以上版本和Xcode 26 beta。<br>    <br>- 可以在macOS 15上运行，但某些高级特性无法使用。<br><br>框架亮点：<br><br>1. 支持管理OCI镜像：帮助开发者管理容器镜像，方便容器操作。<br>    <br>2. 容器化进程：通过创建轻量级虚拟机来运行Linux容器，这些容器能快速启动并且具有低资源占用。<br>    <br>3. 集成vminitd系统：作为虚拟机中的初始进程，vminitd提供了运行时配置、进程启动等功能，并且支持通过gRPC API与容器化的进程交互。<br>    <br>4. 使用Rosetta 2运行Linux/amd64容器：即便在Apple Silicon上，用户也能运行兼容的amd64架构容器。<br>    <br>5. 非隔离网络模式：在macOS 15及以上版本中，容器间不能直接通信。<br><br>开发者可以使用cctl工具测试和试用容器化框架的各项功能，包括OCI镜像管理、容器日志、Linux容器运行等功能。可以参考文档了解详细的API接口。<br><br>感兴趣的小伙伴可以点击：github.com/apple/containerization<img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i2bi29v6h5j31iq0zmqo9.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

苹果推出Containerization框架，允许在Apple Silicon Mac上运行Linux容器。该方案基于Virtualization.framework，支持OCI镜像管理、轻量级虚拟机容器（快速启动/低资源占用）、vminitd系统集成及Rosetta 2运行amd64容器。需macOS 15+Xcode 26 beta开发，部分功能受限。容器间网络隔离，开发者可通过cctl工具测试功能。项目已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-11T14:04:29Z
- **目录日期**: 2025-06-11
