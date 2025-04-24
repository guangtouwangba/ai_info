# #Win11惊现GTA漏洞##GTA圣安地列斯Bug隐藏20年#侠盗猎车圣安地列斯（GTA San Andreas），20年过去了还有新Bug曝出？逆向工程师Adrian揭露，2005年发布的GTA圣安...

**URL**: https://weibo.com/6105753431/PoKZBuKjU

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Win11%E6%83%8A%E7%8E%B0GTA%E6%BC%8F%E6%B4%9E%23&amp;extparam=%23Win11%E6%83%8A%E7%8E%B0GTA%E6%BC%8F%E6%B4%9E%23" data-hide=""><span class="surl-text">#Win11惊现GTA漏洞#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23GTA%E5%9C%A3%E5%AE%89%E5%9C%B0%E5%88%97%E6%96%AFBug%E9%9A%90%E8%97%8F20%E5%B9%B4%23&amp;extparam=%23GTA%E5%9C%A3%E5%AE%89%E5%9C%B0%E5%88%97%E6%96%AFBug%E9%9A%90%E8%97%8F20%E5%B9%B4%23" data-hide=""><span class="surl-text">#GTA圣安地列斯Bug隐藏20年#</span></a><br><br>侠盗猎车圣安地列斯（GTA San Andreas），20年过去了还有新Bug曝出？<br><br>逆向工程师Adrian揭露，2005年发布的GTA圣安地列斯，在Win 11的24H2版本上，遇到了一个诡异的情况——<br><br>Skimmer水上飞机【图1】消失了，到处都找不到，就连秘籍也无法生成！【图2】<br><br>为了调查这个奇怪事件，作者回溯到Windows 11前一个版本23H2，发现Skimmer飞机能够正常出现。<br><br>奇怪的是，尽管GTA的论坛资源充足，却没人提到过这个问题。<br><br>正当他疑惑时，无意中在一个老旧的GitHub页面上，找到了一个名为SilentPatch的修复工具。<br><br>这个工具早在几年前就发布了，曾经解决了不少玩家遇到的问题，似乎能帮助他修复丢失的飞机。<br><br>安装这个补丁后，他再次启动游戏，问题反而更加严重了：整个游戏世界都变得奇怪起来。<br><br>街道和建筑开始错乱，天空的颜色变得很诡异，甚至角色也不再听从指挥——按下的键位会完全失效。【图3】<br><br>整个游戏的画面变得扭曲、混乱，像是被某种无形的力量操控着。<br><br>问题到底出在哪儿呢？<br><br>Adrian查看底层代码发现，在进入Skimmer飞机时，游戏中的CPlane::PreRender函数出现了无限循环，导致游戏冻结。<br><br>具体来说，飞机的叶片角度计算，出现了异常大的值（3.73340132e+29），这个数字大到让后续的计算崩溃了。<br><br>他通过调试发现，Skimmer飞机的叶片速度，来自一个和飞机高度有关的公式，而这个高度值异常高，导致飞机飞得超级高。【图4】<br><br>更离谱的是，Skimmer飞机的定义文件中，20年一直存在错误。<br><br>原本，Skimmer的定义是一艘船（毕竟能在水上游），并没有包含飞机所须的参数。<br><br>然而，在进入游戏时，Skimmer的类型被开发人员强行改为“飞机”。<br><br>真相终于大白——<br><br>Skimmer缺失的轮子参数，导致内存未初始化，造成了后续的错误。<br><br>这个Bug幸存了20年，在Windows 10等此前版本中，Skimmer“前后轮比例”参数（0.7）恰好残留在内存中【图5】，意外让Skimmer勉强正常运行。<br><br>但Windows 11 24H2的更新改变了堆栈的布局（涉及 LeaveCriticalSection 函数），覆盖了残留值，这才让问题暴露出来。【图6】<br><br>玩家目前可手动编辑游戏目录下的data\vehicles.ide文件，找到Skimmer行，并添加“前后轮比例”（0.7），即可恢复飞机Bug。<br><br>这场调查提醒我们，即便是再经典的游戏，也会存在Bug，开发者在编写代码时，必须严格验证所有输入数据，避免不确定的行为。<br><br>感兴趣的小伙伴可以阅读原文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fcookieplmonster.github.io%2F2025%2F04%2F23%2Fgta-san-andreas-win11-24h2-bug%2F" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0s0nu5f10j30sg0g07ec.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0s0nvog7cj30sg0g0wlz.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0s0nwizn3j30sg0g0q7x.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0s0nxnsl3j30sg0g0gnc.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0s0nzupd6j30x20gzamc.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0s0o1iqh6j30x20gz7gz.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

逆向工程师Adrian发现，经典游戏《GTA圣安地列斯》在Win11 24H2版本中出现水上飞机Skimmer消失的Bug。调查发现，该飞机20年来一直存在类型定义错误（本应为船却被强制设为飞机），导致内存未初始化。此前因内存残留值（0.7）勉强运行，但Win11 24H2更新改变了堆栈布局，使Bug彻底爆发，引发游戏画面错乱。解决方法为手动修改游戏文件添加轮子参数。该事件揭示了长期隐藏的代码缺陷如何因系统更新而显现。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T17:03:03Z
- **目录日期**: 2025-04-24
