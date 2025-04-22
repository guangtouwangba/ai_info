# #老旧游戏机爆改服务器##任天堂爆改成服务器#老旧游戏机——任天堂Wii极限挑战！它摇身一变，成了托管博客的服务器。爱折腾的Alex通过安装NetBSD操作系统，将这...

**URL**: https://weibo.com/6105753431/PorCAEFXj

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%80%81%E6%97%A7%E6%B8%B8%E6%88%8F%E6%9C%BA%E7%88%86%E6%94%B9%E6%9C%8D%E5%8A%A1%E5%99%A8%23&amp;extparam=%23%E8%80%81%E6%97%A7%E6%B8%B8%E6%88%8F%E6%9C%BA%E7%88%86%E6%94%B9%E6%9C%8D%E5%8A%A1%E5%99%A8%23" data-hide=""><span class="surl-text">#老旧游戏机爆改服务器#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%BB%BB%E5%A4%A9%E5%A0%82%E7%88%86%E6%94%B9%E6%88%90%E6%9C%8D%E5%8A%A1%E5%99%A8%23&amp;extparam=%23%E4%BB%BB%E5%A4%A9%E5%A0%82%E7%88%86%E6%94%B9%E6%88%90%E6%9C%8D%E5%8A%A1%E5%99%A8%23" data-hide=""><span class="surl-text">#任天堂爆改成服务器#</span></a><br><br>老旧游戏机——任天堂Wii极限挑战！它摇身一变，成了托管博客的服务器。<br><br>爱折腾的Alex通过安装NetBSD操作系统，将这台怀旧游戏机重新赋予了新的生命。这不仅是对Wii的一次“重生”，更是一次对老硬件潜力的探索与挑战。<br><br>故事源于Alex在EMF Camp 2024的跳蚤市场上，发现了一台二手Wii【图1】。<br><br>Alex起初想给Wii装点游戏，回味一下昔日的欢乐。然而，下载操作系统时，他意外地发现NetBSD竟然支持Wii这一平台。<br><br>NetBSD系统以稳定、跨平台著称，甚至可以运行在一些非主流硬件上。<br><br>而且NetBSD不仅兼容性好，维护也很到位，并且还有每日更新。这使得Wii成了一个非常有趣的实验平台。<br><br>而Wii的处理器“Broadway”，这款基于IBM PowerPC 750架构的单核CPU，虽然性能上并不出众，但仍足以应对一些简单的任务，如托管一个静态博客网站。<br><br>值得一提的是，PowerPC 750芯片在消费电子产品中较为罕见，但在航天领域却得到了广泛应用。<br><br>例如，著名的“James Webb Space Telescope（詹姆斯·韦布太空望远镜）”就使用了这款处理器。【图2】<br><br>安装过程相对简单。通过软改Wii，Alex利用现有漏洞工具启动了Homebrew Channel，并成功将NetBSD系统安装到Wii上。<br><br>通过SSH，他远程管理了系统，并配置了网络和基本设置。接着，他安装了轻量级的lighttpd服务器来托管博客。<br><br>尽管Wii成功托管了博客，但在并发请求较多时，系统表现有些吃力。<br><br>为了解决这个问题，作者通过引入Caddy作为反向代理，将TLS终止和证书管理的任务转移到Caddy上，从而减轻Wii的负担。<br><br>此外，为了进一步优化，作者关闭了一些不必要的服务（ntpd）以节省内存和CPU资源。<br><br>为了监控系统运行状态，Alex还编写了一个简单的shell脚本，定期收集Wii的资源使用情况，并将其展示在网页上。<br><br>这样一来，他就可以实时查看Wii的性能表现，确保系统的稳定运行。<br><br>作者发现，Wii不仅能完成这项任务，而且它的功耗也低，仅需18W的电力，甚至比很多云主机服务还要便宜。<br><br>最终，Alex发现Wii不仅能完成任务，而且功耗仅为18W，远低于许多云主机服务，且成本更低。<br><br>这场结合老旧硬件与现代操作系统的实验，不仅展示了技术的无限可能，也为“资源受限”环境下的开发提供了新的思路。<br><br>感兴趣的小伙伴可以阅读原文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fblog.infected.systems%2Fposts%2F2025-04-21-this-blog-is-hosted-on-a-nintendo-wii%2F" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i0pn5ceadyj31601k01kx.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0pn5cuvoxj30zk0m8ngt.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0pn5dx2nlj30rs0ijjwm.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0pn5fb28bj30w50k0jz8.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

Alex将一台老旧任天堂Wii游戏机改造成博客服务器。通过安装NetBSD操作系统（该系统以跨平台兼容性著称）并利用Wii的PowerPC 750处理器，他成功运行了轻量级web服务器。虽然处理并发请求时性能有限，但通过使用Caddy反向代理和优化系统服务，最终实现了稳定运行。整机功耗仅18W，比云服务更节能。这个创意项目不仅赋予旧硬件新生命，还展示了在资源受限环境下技术创新的可能性。值得一提的是，Wii采用的处理器与詹姆斯·韦布太空望远镜同属PowerPC架构。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-22T16:02:50Z
- **目录日期**: 2025-04-22
