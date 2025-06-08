# #统一框架打破跨模态干扰##如何解决模态干扰#多模态检索是AI界的一块硬骨头，尤其“跨模态干扰”问题，一直没被妥善解决。快手联合东北大学的研究团队提出了解法...

**URL**: https://weibo.com/6105753431/PvAgHglRK

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E7%BB%9F%E4%B8%80%E6%A1%86%E6%9E%B6%E6%89%93%E7%A0%B4%E8%B7%A8%E6%A8%A1%E6%80%81%E5%B9%B2%E6%89%B0%23&amp;extparam=%23%E7%BB%9F%E4%B8%80%E6%A1%86%E6%9E%B6%E6%89%93%E7%A0%B4%E8%B7%A8%E6%A8%A1%E6%80%81%E5%B9%B2%E6%89%B0%23" data-hide=""><span class="surl-text">#统一框架打破跨模态干扰#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A6%82%E4%BD%95%E8%A7%A3%E5%86%B3%E6%A8%A1%E6%80%81%E5%B9%B2%E6%89%B0%23&amp;extparam=%23%E5%A6%82%E4%BD%95%E8%A7%A3%E5%86%B3%E6%A8%A1%E6%80%81%E5%B9%B2%E6%89%B0%23" data-hide=""><span class="surl-text">#如何解决模态干扰#</span></a><br><br>多模态检索是AI界的一块硬骨头，尤其“跨模态干扰”问题，一直没被妥善解决。快手联合东北大学的研究团队提出了解法，名为UNITE的统一嵌入框架。<br><br>UNITE的核心是“一个嵌入器搞定文本、图像、视频及其组合”。它通过“模态感知对比学习”，解决模态之间的语义干扰问题。UNITE采用MAMCL机制，训练时只对模态一致的负样本进行对比，避免了不同模态“互相抢戏”。<br><br>整个训练分两阶段：<br><br>- 第一阶段：检索适应，建立基本多模态检索能力，引入细粒度视频-文本数据；<br>    <br>- 第二阶段：指令微调，通过复杂任务增强指令理解与扩展能力。<br>    <br>在多个评测任务中，UNITE都表现亮眼：<br><br>- 图文检索超越E5-V、VLM2Vec；<br>    <br>- 视频检索在CaReBench任务中，7B模型刷新SOTA；<br>    <br>- 指令检索中，UNITE 7B跑赢了更大模型如mmE5 11B和IDMR 26B；<br>    <br>- 通用性方面，在Flickr30K、MSR-VTT等任务中表现稳健。<br>    <br>作者还发现：<br><br>1. 视频-文本数据具备“统一模态”潜力，图文任务中也能吊打图像训练模型；<br>    <br>2. 指令任务更吃Text–Text和Text–Image数据；<br>    <br>3. 细粒度Text-Video样本加入第一阶段训练效果最佳。<br>    <br>更多细节可见论文与代码仓库：<br><br>论文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fpdf%2F2505.19650" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>项目：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Ffriedrichor.github.io%2Fprojects%2FUNITE" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a> <a href="https://weibo.com/ttarticle/p/show?id=2309405175275876122799" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">打破跨模态干扰，快手东北大学联合提出统一多模态框架</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3ly1i27x9fbocbj30rs0fm41a.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

快手与东北大学团队提出统一嵌入框架UNITE，通过模态感知对比学习(MAMCL)解决跨模态干扰问题。该框架使用单一嵌入器处理文本、图像、视频及其组合，采用两阶段训练：检索适应建立基础能力，指令微调增强复杂任务处理。实验显示，UNITE在图文检索、视频检索(CaReBench SOTA)和指令任务中超越E5-V等模型，7B版本性能优于更大模型。研究发现视频-文本数据具有统一模态潜力，细粒度样本对第一阶段训练效果最佳。论文和代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-08T16:01:43Z
- **目录日期**: 2025-06-08
