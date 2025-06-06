# #80%token在训练中起副作用#大模型强化学习超越二八法则#近期arxiv最热门论文，Qwen&清华LeapLab团队最新成果：在强化学习训练大模型推理能力时，仅仅20%的高熵t...

**URL**: https://weibo.com/6105753431/Pvg4K17Pm

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%2380%25token%E5%9C%A8%E8%AE%AD%E7%BB%83%E4%B8%AD%E8%B5%B7%E5%89%AF%E4%BD%9C%E7%94%A8%23&amp;extparam=%2380%25token%E5%9C%A8%E8%AE%AD%E7%BB%83%E4%B8%AD%E8%B5%B7%E5%89%AF%E4%BD%9C%E7%94%A8%23" data-hide=""><span class="surl-text">#80%token在训练中起副作用#</span></a>大模型强化学习超越二八法则#<br><br>近期arxiv最热门论文，Qwen&amp;清华LeapLab团队最新成果：<br><br>在强化学习训练大模型推理能力时，仅仅20%的高熵token就能撑起整个训练效果，甚至比用全部token训练还要好。【图1】<br><br>团队用这个发现在Qwen3-32B上创造了新的SOTA记录：AIME’24上达到63.5分，AIME’25上达到56.7分。【图2】<br><br>这是600B参数以下直接从base模型训练的最高分。<br><br>最大响应长度从20k延长到29k，AIME’24的分数更是飙升到了68.1分。<br><br>经典的二八法则（或帕累托法则）指出，通常80%的结果由20%的关键因素驱动，但剩下80%也是不能轻易舍弃的。<br><br>但是在大模型强化学习这里，80%低熵token不仅可以舍弃，甚至还可能起副作用，所以这篇论文被命名为“超越二八法则”。【图3】<br><br>此外，团队还从token熵的角度探究了RL对LLM的主要影响，并进一步讨论了RL与SFT的区别、LLM RL的特殊性与clip-higher相较于entropy bonus的优势。<br><br>点击链接，一起深入理解这项研究：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2F8VNXnvjoapEdHlFfxHsWfQ" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">Qwen&amp;清华团队颠覆常识：大模型强化学习仅用20%关键token，比用全部token训练还好</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i25g6n13q0j30zk07s77r.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i25g6p0bdcj30yu0e8ds7.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i25g6r0i2xj30vy0acdkl.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

Qwen和清华LeapLab团队的最新研究发现，在大模型强化学习训练中，仅20%的高熵token就足以支撑训练效果，甚至优于使用全部token。这一发现超越了传统的二八法则，因为剩余的80%低熵token不仅可舍弃，还可能产生副作用。团队在Qwen3-32B模型上实现了新SOTA成绩，AIME’24得分达63.5分，最大响应长度延长至29k后分数提升至68.1分。研究还探讨了RL对LLM的影响、RL与SFT的区别，以及clip-higher相比entropy bonus的优势。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-06T06:02:23Z
- **目录日期**: 2025-06-06
