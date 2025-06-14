# #80%token在训练中起副作用##大模型强化学习超越二八法则#近期arxiv最热门论文，Qwen&清华LeapLab团队最新成果：在强化学习训练大模型推理能力时，仅仅20%的高熵...

**URL**: https://weibo.com/6105753431/Pvg4K17Pm

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%2380%25token%E5%9C%A8%E8%AE%AD%E7%BB%83%E4%B8%AD%E8%B5%B7%E5%89%AF%E4%BD%9C%E7%94%A8%23&amp;extparam=%2380%25token%E5%9C%A8%E8%AE%AD%E7%BB%83%E4%B8%AD%E8%B5%B7%E5%89%AF%E4%BD%9C%E7%94%A8%23" data-hide=""><span class="surl-text">#80%token在训练中起副作用#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%E5%BC%BA%E5%8C%96%E5%AD%A6%E4%B9%A0%E8%B6%85%E8%B6%8A%E4%BA%8C%E5%85%AB%E6%B3%95%E5%88%99%23&amp;extparam=%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%E5%BC%BA%E5%8C%96%E5%AD%A6%E4%B9%A0%E8%B6%85%E8%B6%8A%E4%BA%8C%E5%85%AB%E6%B3%95%E5%88%99%23" data-hide=""><span class="surl-text">#大模型强化学习超越二八法则#</span></a><br><br>近期arxiv最热门论文，Qwen&amp;清华LeapLab团队最新成果：<br><br>在强化学习训练大模型推理能力时，仅仅20%的高熵token就能撑起整个训练效果，甚至比用全部token训练还要好。【图1】<br><br>团队用这个发现在Qwen3-32B上创造了新的SOTA记录：AIME’24上达到63.5分，AIME’25上达到56.7分。【图2】<br><br>这是600B参数以下直接从base模型训练的最高分。<br><br>最大响应长度从20k延长到29k，AIME’24的分数更是飙升到了68.1分。<br><br>经典的二八法则（或帕累托法则）指出，通常80%的结果由20%的关键因素驱动，但剩下80%也是不能轻易舍弃的。<br><br>但是在大模型强化学习这里，80%低熵token不仅可以舍弃，甚至还可能起副作用，所以这篇论文被命名为“超越二八法则”。【图3】<br><br>此外，团队还从token熵的角度探究了RL对LLM的主要影响，并进一步讨论了RL与SFT的区别、LLM RL的特殊性与clip-higher相较于entropy bonus的优势。<br><br>点击链接，一起深入理解这项研究：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2F8VNXnvjoapEdHlFfxHsWfQ" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">Qwen&amp;清华团队颠覆常识：大模型强化学习仅用20%关键token，比用全部token训练还好</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i25g6n13q0j30zk07s77r.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i25g6p0bdcj30yu0e8ds7.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i25g6r0i2xj30vy0acdkl.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

Qwen和清华LeapLab团队最新研究发现，在大模型强化学习中，仅20%的高熵token就能主导训练效果，甚至优于使用全部token。这一发现颠覆了传统二八法则，表明80%低熵token不仅可舍弃，还可能产生副作用。团队基于此在Qwen3-32B模型上取得突破：AIME'24得分63.5，AIME'25得分56.7，创下600B参数以下base模型的最高记录。研究还探讨了强化学习与监督微调的区别，以及clip-higher方法的优势。成果显示大模型训练存在更高效的数据选择策略。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-06T08:04:47Z
- **目录日期**: 2025-06-06
