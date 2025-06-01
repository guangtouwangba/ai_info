# #推理提速2倍还省内存？Mamba作者又整新活了。Tri Dao团队刚发布两种专为推理设计的注意力机制：Grouped-Tied Attention（GTA）和Grouped Latent Attention（GLA...

**URL**: https://weibo.com/6105753431/PuwnQDWCX

## 原始摘要

#推理提速2倍还省内存？Mamba作者又整新活了。<br><br>Tri Dao团队刚发布两种专为推理设计的注意力机制：Grouped-Tied Attention（GTA）和Grouped Latent Attention（GLA）。它们在不牺牲模型生成质量的前提下，显著提升了解码速度和内存利用率，尤其适合长上下文场景。<br><br>团队在四种模型规模上测试了GTA和GLA，指标涵盖质量（困惑度、7个下游任务）与效率（解码延迟、吞吐量、KV缓存量）。结果显示：<br><br>- GTA在中大型模型中质量优于GQA；<br><br>- GLA与MLA在精度上相当，但效率更高；<br><br>- GLA在预填充长度32K、64K时的吞吐量明显领先；<br><br>- 并发处理能力也更强，能更好应对长文本和不均衡负载。<br> <a href="https://weibo.com/ttarticle/p/show?id=2309405172743271481520" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">Mamba核心作者新作：取代DeepSeek在用的注意力机制，专为推理打造</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3ly1i1zugwccmmj30qp0f1jsu.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

Tri Dao团队推出两种新型注意力机制GTA和GLA，专为优化推理设计。实验表明，GTA在中大型模型中的生成质量优于GQA，而GLA与MLA精度相当但效率更高。两种机制显著提升了长上下文场景下的解码速度、吞吐量和内存利用率，尤其擅长处理32K/64K长文本和不均衡负载，在保持模型质量的同时实现推理加速和资源节省。研究覆盖四种模型规模，测试包括困惑度、下游任务表现及多项效率指标。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-01T08:02:53Z
- **目录日期**: 2025-06-01
