# #让DeepSeek专家网络动起来可降低推理延迟##华为让MoE专家动起来推理延迟降 10%# 要问最近哪个模型最火，混合专家模型（MoE，Mixture of Experts）绝对是榜上提...

**URL**: https://weibo.com/6105753431/PsGHei35J

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%AE%A9DeepSeek%E4%B8%93%E5%AE%B6%E7%BD%91%E7%BB%9C%E5%8A%A8%E8%B5%B7%E6%9D%A5%E5%8F%AF%E9%99%8D%E4%BD%8E%E6%8E%A8%E7%90%86%E5%BB%B6%E8%BF%9F%23&amp;extparam=%23%E8%AE%A9DeepSeek%E4%B8%93%E5%AE%B6%E7%BD%91%E7%BB%9C%E5%8A%A8%E8%B5%B7%E6%9D%A5%E5%8F%AF%E9%99%8D%E4%BD%8E%E6%8E%A8%E7%90%86%E5%BB%B6%E8%BF%9F%23" data-hide=""><span class="surl-text">#让DeepSeek专家网络动起来可降低推理延迟#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%8D%8E%E4%B8%BA%E8%AE%A9MoE%E4%B8%93%E5%AE%B6%E5%8A%A8%E8%B5%B7%E6%9D%A5%E6%8E%A8%E7%90%86%E5%BB%B6%E8%BF%9F%E9%99%8D+10%25%23&amp;extparam=%23%E5%8D%8E%E4%B8%BA%E8%AE%A9MoE%E4%B8%93%E5%AE%B6%E5%8A%A8%E8%B5%B7%E6%9D%A5%E6%8E%A8%E7%90%86%E5%BB%B6%E8%BF%9F%E9%99%8D+10%25%23" data-hide=""><span class="surl-text">#华为让MoE专家动起来推理延迟降 10%#</span></a> <br><br>要问最近哪个模型最火，混合专家模型（MoE，Mixture of Experts）绝对是榜上提名的那一个。<br><br>它的巧妙之处，就在于把不同的任务分配给擅长处理的专家网络，让整个系统性能得以提升。<br><br>但你知道吗？<br><br>正是这个关键的专家网络，也是严重影响系统推理性能的因素之一。<br><br>因为在大量任务来临之际（尤其是超大规模时），MoE并不是以“雨露均沾”的方式去分配——专家网络们的负载均衡问题，就会显得尤为突出。<br><br>这个问题的根源，是因为某些专家网络总是被频繁调用（热专家），而另一些专家网络则鲜有机会派上用场（冷专家）。<br><br>没错，MoE里的“专家们”也是有冷热之分的，而且被调用频率的差距甚至可以达到一个数量级以上！<br><br>如此负载不均衡的现象，就会导致整个系统推理的时间被延长，以及还有资源利用率、系统性能受限等问题。<br><br>那么此局又该如何破解？<br><br>别急，华为团队已经给出了一种有效解法，直接让DeepSeek-V3在理论上的推理延迟可降低约10%、吞吐量可提升约10%。<br><br>值得一提的是，团队还将在近期准备把这个解法全面开源了；那么接下来，我们就来深入了解一下。<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FUN9-sORx7hgNgb6tUrxz4A" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">华为：让DeepSeek的“专家们”动起来，推理延迟降10%！</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3ly1i1lwkavnxsj30u00gok37.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

华为团队针对混合专家模型(MoE)中专家网络负载不均衡问题提出创新解决方案，通过动态调整专家调用策略，显著改善了系统性能。该方案使DeepSeek-V3模型的推理延迟降低约10%，吞吐量提升约10%。MoE模型中存在"热专家"被频繁调用而"冷专家"闲置的问题，导致资源利用率低下。华为的方法有效平衡了专家网络负载，并将开源这一解决方案。这一突破对提升大规模AI系统的推理效率具有重要意义。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-20T07:03:41Z
- **目录日期**: 2025-05-20
