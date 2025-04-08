# #Git诞生20年##Git背后的故事#今天，Git正式满20岁了！你可能想不到，全球程序员每天离不开的Git，其实是Linus Torvalds一拍脑门写出来的“临时方案”。时间回到...

**URL**: https://weibo.com/6105753431/PmheVFnqb

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Git%E8%AF%9E%E7%94%9F20%E5%B9%B4%23&amp;extparam=%23Git%E8%AF%9E%E7%94%9F20%E5%B9%B4%23" data-hide=""><span class="surl-text">#Git诞生20年#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Git%E8%83%8C%E5%90%8E%E7%9A%84%E6%95%85%E4%BA%8B%23&amp;extparam=%23Git%E8%83%8C%E5%90%8E%E7%9A%84%E6%95%85%E4%BA%8B%23" data-hide=""><span class="surl-text">#Git背后的故事#</span></a><br><br>今天，Git正式满20岁了！<br><br>你可能想不到，全球程序员每天离不开的Git，其实是Linus Torvalds一拍脑门写出来的“临时方案”。<br><br>时间回到2005年，那时Linux内核团队用的是一家商业公司的版本控制工具——BitKeeper。<br><br>但因为授权纠纷，双方最终闹掰，Linux社区一时间竟然无工具可用。<br><br>面对这突如其来的窘境，开发者Linus一气之下，自己造了个版本控制系统。<br><br>该系统的设计初衷并非为了团队协作开发，而只是为了更高效地发布补丁和压缩包。<br><br>Linus的态度也非常佛系：“我就写个引擎，界面你们自己搞。”<br><br>所以，最初的Git只有一堆底层工具，如update-cache、write-tree、commit-tree等。<br><br>但也正因为Git采用了去中心化架构，每个人本地都有一个完整的仓库，所有历史记录清晰透明，这种机制一下就碾压了当时的大多数版本控制工具。<br><br>于是，全球开发者们开始了“二次创作”，逐步构建起更完整的Git工具链，像我们今天熟悉的git log、git rebase等指令，都是后来慢慢发展起来的。<br><br>还有个有趣的插曲：Git在诞生几个月后，竟然被一家公司用来“发广告”。<br><br>这家公司要在全美几百个商场里的广告屏幕投放不同版本的素材，这些素材组合多、改动频繁，传统方法完全hold不住。<br><br>最后他们灵机一动，决定用Git做内容分发：每个屏幕对应一个分支，素材更新就commit，远程pull再checkout，完美解决。<br><br>这事还催生了Git的大更新，比如支持SSL、HTTP并发传输、断点续传……这些奇奇怪怪的需求，反而让Git更强大了。<br><br>现在20年过去了，Git几乎成了所有开发者的“标配”。无论是OpenAI、Meta，还是Google、Microsoft，写代码基本都离不开它。<br><br>你在VS Code里点“提交”的时候，其实底层运行的，还是Linus当年灵机一动写下的命令。<br><br>更神奇的是，Git那套用SHA-1哈希指向不可变对象树的底层架构，20年来几乎没变过。<br><br>即使是今天这些大模型项目，也依旧在用这套系统。<br><br>所以，20岁生日快乐，Git。<br><br>感兴趣的小伙伴可以点击：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fblog.gitbutler.com%2F20-years-of-git%2F" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i095vw6fdmj30m80fvah5.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i095vxg6xxj31ja18qhdt.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

Git诞生20周年，最初是Linux创始人Linus Torvalds在2005年因BitKeeper授权纠纷而临时开发的版本控制系统。其去中心化架构让每个用户拥有完整仓库，迅速超越同期工具。开发者社区逐步完善了git log等常用指令。有趣的是，Git早期曾被广告公司用于商场屏幕内容分发，这促使它新增了SSL支持等功能。20年来核心架构未变，现已成为全球开发标配，支撑着从Linux到AI项目的代码管理。这个源于"临时方案"的工具，最终彻底改变了软件开发协作方式。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-08T05:03:10Z
- **目录日期**: 2025-04-08
