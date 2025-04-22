# #微软抄袭个人开发者##开源项目作者被微软欺负#微软抄袭个人开发者的项目，连句出处都没标？事情起因是，开源项目Spegel的作者Philip Laine，控诉微软未经正当署...

**URL**: https://weibo.com/6105753431/PopiUgJBX

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%BE%AE%E8%BD%AF%E6%8A%84%E8%A2%AD%E4%B8%AA%E4%BA%BA%E5%BC%80%E5%8F%91%E8%80%85%23&amp;extparam=%23%E5%BE%AE%E8%BD%AF%E6%8A%84%E8%A2%AD%E4%B8%AA%E4%BA%BA%E5%BC%80%E5%8F%91%E8%80%85%23" data-hide=""><span class="surl-text">#微软抄袭个人开发者#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%BC%80%E6%BA%90%E9%A1%B9%E7%9B%AE%E4%BD%9C%E8%80%85%E8%A2%AB%E5%BE%AE%E8%BD%AF%E6%AC%BA%E8%B4%9F%23&amp;extparam=%23%E5%BC%80%E6%BA%90%E9%A1%B9%E7%9B%AE%E4%BD%9C%E8%80%85%E8%A2%AB%E5%BE%AE%E8%BD%AF%E6%AC%BA%E8%B4%9F%23" data-hide=""><span class="surl-text">#开源项目作者被微软欺负#</span></a><br><br>微软抄袭个人开发者的项目，连句出处都没标？<br><br>事情起因是，开源项目Spegel的作者Philip Laine，控诉微软未经正当署名，直接将其项目“改头换面”推出新产品，试图据为己有。<br><br>消息一出，开源社区一片哗然，微软一时间被“千夫所指”。<br><br>Spegel这个项目，是一个面向Kubernetes的P2P镜像分发工具，解决的是镜像仓库宕机导致集群无法扩容的问题。<br><br>该项目一经开源便受到开发者追捧，迅速揽获了超过1700个Star。【图1】<br><br>项目走红后，微软主动联系Philip，表示希望探讨合作。<br><br>Philip这边表示，他很愿意配合，包括解答架构问题，协助部署等，表现出极大诚意。<br><br>然而，微软在一顿“问问问”后，也不联系他，没下文了。<br><br>直到今年巴黎KubeCon大会上，他在一场讲Kubernetes镜像分发的演讲中，居然看到Spegel被提为“灵感来源”，而这是微软的一个全新项目：Peerd，功能定位几乎一模一样。<br><br>他一看源码，不仅项目结构雷同，这些还都极其相似：<br><br>- 复制了Spegel中的函数定义、注释和测试用例；【图2】<br>- 测试代码里还残留着他老东家的名字；【图3】<br>- 关键代码几乎原封不动搬了过去；<br><br>问题是，该项目没有明确标明来源，仅在README结尾，“感谢”了一句原作署名和代码出处。【图4】<br><br>也就是说，在没有贡献代码、没有协作开发、没有标注来源的情况下，微软拿了Spegel的代码，直接做了自己的产品，还打着微软名号发布。<br><br>这让Philip很是沮丧，一度怀疑Spegel还有没有继续维护下去的意义。<br><br>但最终，他选择坚持开发完这个项目。<br><br>面对舆论压力，微软随后提交了PR，在相关文件头部添加了署名，并发布声明称“我们听见了社区反馈”、“感谢原作者”，同时解释Peerd是为了支持artifact streaming，而Spegel并未涵盖这一范围。<br><br>但社区并不买账，网友表示，既然项目加入了MIT协议，基本的署名和尊重不能没有。<br><br>MIT协议建立在对开源精神的信任之上，若被大厂滥用，开源反而会沦为“免费素材库”。<br><br>有网友用一句话总结道：开源是社区的财富，不该成为“灵感掠夺”的温床。<br><br>原文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fphiliplaine.com%2Fposts%2Fgetting-forked-by-microsoft%2F" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>项目：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fspegel-org%2Fspegel" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0pcwyi53tj30zk0mcds5.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0pcwztzftj312m17kdtj.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i0pcx2fov3j30ik059gnm.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0pcx3vapaj30o3032wfq.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

微软被指控抄袭个人开发者Philip Laine的开源项目Spegel（Kubernetes镜像分发工具），将其代码几乎原样复制到新项目Peerd中，仅在不显眼处标注"灵感来源"。开发者指出微软复制了函数定义、注释、测试代码甚至残留的老东家名称。尽管项目采用MIT协议允许使用，但社区批评微软未给予基本署名尊重。微软事后补加署名并辩解Peerd扩展了新功能，但开发者认为这违背开源精神。事件引发对开源项目被大厂滥用的担忧，强调开源应保护创作者权益而非沦为免费素材库。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-22T03:16:59Z
- **目录日期**: 2025-04-22
