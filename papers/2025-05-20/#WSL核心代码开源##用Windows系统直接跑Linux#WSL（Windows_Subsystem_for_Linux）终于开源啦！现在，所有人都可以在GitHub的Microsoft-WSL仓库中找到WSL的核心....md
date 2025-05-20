# #WSL核心代码开源##用Windows系统直接跑Linux#WSL（Windows Subsystem for Linux）终于开源啦！现在，所有人都可以在GitHub的Microsoft/WSL仓库中找到WSL的核心...

**URL**: https://weibo.com/6105753431/PsFZapfqA

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23WSL%E6%A0%B8%E5%BF%83%E4%BB%A3%E7%A0%81%E5%BC%80%E6%BA%90%23&amp;extparam=%23WSL%E6%A0%B8%E5%BF%83%E4%BB%A3%E7%A0%81%E5%BC%80%E6%BA%90%23" data-hide=""><span class="surl-text">#WSL核心代码开源#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E7%94%A8Windows%E7%B3%BB%E7%BB%9F%E7%9B%B4%E6%8E%A5%E8%B7%91Linux%23&amp;extparam=%23%E7%94%A8Windows%E7%B3%BB%E7%BB%9F%E7%9B%B4%E6%8E%A5%E8%B7%91Linux%23" data-hide=""><span class="surl-text">#用Windows系统直接跑Linux#</span></a><br><br>WSL（Windows Subsystem for Linux）终于开源啦！<br><br>现在，所有人都可以在GitHub的Microsoft/WSL仓库中找到WSL的核心代码，不仅可以自由构建定制版本，甚至还能深度参与WSL的持续开发与优化。<br><br>给不熟悉的朋友们介绍一下WSL是什么，作为微软官方开发的一项功能，它允许用户在Windows系统上直接运行Linux环境。<br><br>它由一系列分布式组件构成，部分运行于Windows，部分运行于WSL 2虚拟机中。具有如【图2】所示的架构。<br><br>此次开源的组件还包括此前已公开的部分：<br>- microsoft/wslg：为WSL提供Wayland和X服务器相关功能的支持<br>- microsoft/WSL2-Linux-Kernel：WSL2所使用的Linux内核源码<br><br>以下组件仍属于Windows系统镜像，暂未开源：<br>- Lxcore.sys：驱动WSL 1的内核级组件<br>- P9rdr.sys 和 p9np.dll：实现“\wsl.localhost”文件系统重定向（从Windows到Linux）<br><br>自Windows 10周年更新首次发布以来，WSL如今已迈入2.0.0时代，引入了镜像网络、DNS隧道、Session 0支持、代理和防火墙等功能。<br><br>目前，WSL已经更新到了2.5.7版本。<br><br>开源仓库：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fmicrosoft%2FWSL%2Freleases%2Ftag%2F2.5.7" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1ltfee8wjj318k1iswyi.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1ltffmdz5j30rs0pydje.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

微软宣布开源WSL（Windows Subsystem for Linux）核心代码，用户现可通过GitHub获取并参与开发。WSL允许在Windows系统直接运行Linux环境，由Windows和WSL 2虚拟机组件构成。此次开源包括WSLG（Wayland/X服务器支持）和WSL2-Linux-Kernel，但部分驱动和文件系统组件仍闭源。WSL自2016年发布后已迭代至2.5.7版本，新增镜像网络、DNS隧道等功能。开发者可自由构建定制版本，推动WSL生态发展。开源地址：github.com/microsoft/WSL。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-20T06:04:01Z
- **目录日期**: 2025-05-20
