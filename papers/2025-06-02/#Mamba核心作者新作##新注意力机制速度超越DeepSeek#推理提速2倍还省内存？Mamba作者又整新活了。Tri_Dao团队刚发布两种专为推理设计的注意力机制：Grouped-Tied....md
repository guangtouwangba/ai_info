# #Mamba核心作者新作##新注意力机制速度超越DeepSeek#推理提速2倍还省内存？Mamba作者又整新活了。Tri Dao团队刚发布两种专为推理设计的注意力机制：Grouped-Tied...

**URL**: https://weibo.com/6105753431/PuwnQDWCX

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Mamba%E6%A0%B8%E5%BF%83%E4%BD%9C%E8%80%85%E6%96%B0%E4%BD%9C%23&amp;extparam=%23Mamba%E6%A0%B8%E5%BF%83%E4%BD%9C%E8%80%85%E6%96%B0%E4%BD%9C%23" data-hide=""><span class="surl-text">#Mamba核心作者新作#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%96%B0%E6%B3%A8%E6%84%8F%E5%8A%9B%E6%9C%BA%E5%88%B6%E9%80%9F%E5%BA%A6%E8%B6%85%E8%B6%8ADeepSeek%23&amp;extparam=%23%E6%96%B0%E6%B3%A8%E6%84%8F%E5%8A%9B%E6%9C%BA%E5%88%B6%E9%80%9F%E5%BA%A6%E8%B6%85%E8%B6%8ADeepSeek%23" data-hide=""><span class="surl-text">#新注意力机制速度超越DeepSeek#</span></a><br><br>推理提速2倍还省内存？Mamba作者又整新活了。<br><br>Tri Dao团队刚发布两种专为推理设计的注意力机制：Grouped-Tied Attention（GTA）和Grouped Latent Attention（GLA）。它们在不牺牲模型生成质量的前提下，显著提升了解码速度和内存利用率，尤其适合长上下文场景。<br><br>团队在四种模型规模上测试了GTA和GLA，指标涵盖质量（困惑度、7个下游任务）与效率（解码延迟、吞吐量、KV缓存量）。结果显示：<br><br>- GTA在中大型模型中质量优于GQA；<br><br>- GLA与MLA在精度上相当，但效率更高；<br><br>- GLA在预填充长度32K、64K时的吞吐量明显领先；<br><br>- 并发处理能力也更强，能更好应对长文本和不均衡负载。<br> <a href="https://weibo.com/ttarticle/p/show?id=2309405172743271481520" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">Mamba核心作者新作：取代DeepSeek在用的注意力机制，专为推理打造</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3ly1i1zugwccmmj30qp0f1jsu.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

Mamba核心作者Tri Dao团队最新发布两种高效注意力机制GTA和GLA，专为优化推理性能设计。实验显示，这两种机制在保持模型生成质量的同时，显著提升解码速度（最高提速2倍）并降低内存消耗，尤其擅长处理长文本。GTA在中大型模型中表现优于GQA，GLA则与MLA精度相当但效率更高，在32K/64K长文本处理时吞吐量优势明显，且具备更强的并发处理能力。新机制有望替代现有方案，为长上下文场景提供更高效的推理解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-02T01:30:45Z
- **目录日期**: 2025-06-02
