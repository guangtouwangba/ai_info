# #AR眼镜加手机搞定编程##不用电脑也能编程办公#没有电脑，让你进行两周的编程，你能做到吗？还真有出门旅行不想带电脑的国外网友做到了，他的装备很简单：二手Pi...

**URL**: https://weibo.com/6105753431/PsxQYcrHk

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AR%E7%9C%BC%E9%95%9C%E5%8A%A0%E6%89%8B%E6%9C%BA%E6%90%9E%E5%AE%9A%E7%BC%96%E7%A8%8B%23&amp;extparam=%23AR%E7%9C%BC%E9%95%9C%E5%8A%A0%E6%89%8B%E6%9C%BA%E6%90%9E%E5%AE%9A%E7%BC%96%E7%A8%8B%23" data-hide=""><span class="surl-text">#AR眼镜加手机搞定编程#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%8D%E7%94%A8%E7%94%B5%E8%84%91%E4%B9%9F%E8%83%BD%E7%BC%96%E7%A8%8B%E5%8A%9E%E5%85%AC%23&amp;extparam=%23%E4%B8%8D%E7%94%A8%E7%94%B5%E8%84%91%E4%B9%9F%E8%83%BD%E7%BC%96%E7%A8%8B%E5%8A%9E%E5%85%AC%23" data-hide=""><span class="surl-text">#不用电脑也能编程办公#</span></a><br><br>没有电脑，让你进行两周的编程，你能做到吗？<br><br>还真有出门旅行不想带电脑的国外网友做到了，他的装备很简单：二手Pixel 8 Pro、二手Xreal Air 2 Pro AR眼镜和Samers折叠键盘，总共花费636美元。【图1】<br><br>在编译Nim源码的测试中，这台手机做出了不错的成绩：耗时11分20秒，优于Thinkpad T450s（14分20秒），但不及Framework 13（4分15秒）。<br><br>那么，这样编程的实际体感如何呢？这名网友表示：有好有坏。<br><br>好的一方面是，这一套能揣兜里带走的编程设施提供了笔记本电脑无法提供的自由感。<br><br>麻烦的一方面也有，不过也能克服，稍微总结一下就是：<br><br>一、Linux环境<br><br>在安卓系统上搭建Linux环境颇为复杂。常见方法有四种：【图2】<br><br>1. 模拟x86_64的虚拟机<br><br>2. Termux：一款Android应用，提供终端模拟器、轻量级Linux用户环境和一组能在该环境中运行的软件包。<br><br>3. 在chroot中运行arm64二进制文件：本质上是一个隔离的目录，程序在其中运行，与文件系统的其他部分隔离。但需要root权限。<br><br>4. proot：与chroot类似，但不需要root权限。<br><br>经过大量实验，网友发现：虚拟机和proot太慢太笨重，Termux受限于安卓的Bionic C库，而chroot几乎没有性能损失，任何能编译为arm64的程序似乎都能运行。<br><br>至于选择什么样的Linux 发行版，网友考虑的几个关键点如下：轻量、支持aarch64、不使用systemd、有chroot支持、使用glibc。<br><br>最终，Void Linux的aarch64 glibc根文件系统压缩包脱颖而出，完美符合需求，运行得非常流畅。<br><br>二、  AR眼镜<br><br>这台眼镜支持电致变色调光，可以调暗镜片以减少环境光，显示效果非常棒。【图3】<br><br>最大的缺点是视野太大，需要转动眼球才能看清屏幕边缘，且难以对焦。不过通过在窗口管理器中为屏幕上下添加了一些额外边距，这个问题也得到了缓解。<br><br>除此之外，安卓的多显示器模式体验不佳，但可以将手机分辨率调整为1080p并镜像到眼镜，效果很好。<br><br>看完网友的分享，不由得畅想一下，随着AR眼镜的改进和Linux的持续灵活性，在户外边玩边编程好像也不是梦了？<br><br>感兴趣的朋友可以看看博客原文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fholdtherobot.com%2Fblog%2F2025%2F05%2F11%2Flinux-on-android-with-ar-glasses%2F" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1kti866ulj30zk0m01ay.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1ktiepqv1j30zk0na1kx.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1ktifnqn0j30xf0zktrl.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

国外网友使用二手Pixel 8 Pro手机、Xreal Air 2 Pro AR眼镜和折叠键盘（总价636美元）实现了移动编程。通过Void Linux的aarch64 glibc根文件系统在安卓搭建Linux环境，编译速度优于部分笔记本。AR眼镜显示效果优秀但视野过大需调整，手机1080p镜像显示效果良好。该方案提供了便携性优势，虽然存在环境搭建复杂、视觉适配等问题，但展示了移动设备结合AR技术进行编程办公的可行性，为户外工作提供了新思路。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T08:03:44Z
- **目录日期**: 2025-05-19
