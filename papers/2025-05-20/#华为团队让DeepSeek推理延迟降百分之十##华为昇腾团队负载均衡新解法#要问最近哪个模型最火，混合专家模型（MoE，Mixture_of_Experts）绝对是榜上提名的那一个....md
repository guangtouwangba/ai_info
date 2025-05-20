# #华为团队让DeepSeek推理延迟降百分之十##华为昇腾团队负载均衡新解法#要问最近哪个模型最火，混合专家模型（MoE，Mixture of Experts）绝对是榜上提名的那一个...

**URL**: https://weibo.com/6105753431/PsHbw9f9k

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%8D%8E%E4%B8%BA%E5%9B%A2%E9%98%9F%E8%AE%A9DeepSeek%E6%8E%A8%E7%90%86%E5%BB%B6%E8%BF%9F%E9%99%8D%E7%99%BE%E5%88%86%E4%B9%8B%E5%8D%81%23&amp;extparam=%23%E5%8D%8E%E4%B8%BA%E5%9B%A2%E9%98%9F%E8%AE%A9DeepSeek%E6%8E%A8%E7%90%86%E5%BB%B6%E8%BF%9F%E9%99%8D%E7%99%BE%E5%88%86%E4%B9%8B%E5%8D%81%23" data-hide=""><span class="surl-text">#华为团队让DeepSeek推理延迟降百分之十#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%8D%8E%E4%B8%BA%E6%98%87%E8%85%BE%E5%9B%A2%E9%98%9F%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%96%B0%E8%A7%A3%E6%B3%95%23&amp;extparam=%23%E5%8D%8E%E4%B8%BA%E6%98%87%E8%85%BE%E5%9B%A2%E9%98%9F%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%96%B0%E8%A7%A3%E6%B3%95%23" data-hide=""><span class="surl-text">#华为昇腾团队负载均衡新解法#</span></a><br><br>要问最近哪个模型最火，混合专家模型（MoE，Mixture of Experts）绝对是榜上提名的那一个。<br><br>它的巧妙之处，就在于把不同的任务分配给擅长处理的专家网络，让整个系统性能得以提升。<br><br>但你知道吗？<br><br>正是这个关键的专家网络，也是严重影响系统推理性能的因素之一。<br><br>因为在大量任务来临之际（尤其是超大规模时），MoE并不是以“雨露均沾”的方式去分配——专家网络们的负载均衡问题，就会显得尤为突出。<br><br>这个问题的根源，是因为某些专家网络总是被频繁调用（热专家），而另一些专家网络则鲜有机会派上用场（冷专家）。<br><br>没错，MoE里的“专家们”也是有冷热之分的，而且被调用频率的差距甚至可以达到一个数量级以上！<br><br>如此负载不均衡的现象，就会导致整个系统推理的时间被延长，以及还有资源利用率、系统性能受限等问题。<br><br>那么此局又该如何破解？<br><br>别急，华为团队已经给出了一种有效解法，直接让DeepSeek-V3在理论上的推理延迟可降低约10%、吞吐量可提升约10%。<br><br>值得一提的是，团队还将在近期准备把这个解法全面开源了；那么接下来，我们就来深入了解一下。<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FUN9-sORx7hgNgb6tUrxz4A" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">华为：让DeepSeek的“专家们”动起来，推理延迟降10%！</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1lypiahzuj30zk0jrk8o.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

华为团队针对混合专家模型（MoE）中的负载不均衡问题提出新解法，通过优化专家网络调用频率差异（热/冷专家差距达10倍以上），显著提升系统性能。该方法使DeepSeek-V3推理延迟降低约10%、吞吐量提升10%，并计划近期开源。该方案解决了MoE模型中因任务分配不均导致的资源利用率低、推理延迟高等核心问题。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-20T07:02:49Z
- **目录日期**: 2025-05-20
