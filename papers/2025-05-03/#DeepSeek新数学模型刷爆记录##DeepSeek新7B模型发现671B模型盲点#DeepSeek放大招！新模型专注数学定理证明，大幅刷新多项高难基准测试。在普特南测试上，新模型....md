# #DeepSeek新数学模型刷爆记录##DeepSeek新7B模型发现671B模型盲点#DeepSeek放大招！新模型专注数学定理证明，大幅刷新多项高难基准测试。在普特南测试上，新模型...

**URL**: https://weibo.com/6105753431/PpN2nc9Dv

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23DeepSeek%E6%96%B0%E6%95%B0%E5%AD%A6%E6%A8%A1%E5%9E%8B%E5%88%B7%E7%88%86%E8%AE%B0%E5%BD%95%23&amp;extparam=%23DeepSeek%E6%96%B0%E6%95%B0%E5%AD%A6%E6%A8%A1%E5%9E%8B%E5%88%B7%E7%88%86%E8%AE%B0%E5%BD%95%23" data-hide=""><span class="surl-text">#DeepSeek新数学模型刷爆记录#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23DeepSeek%E6%96%B07B%E6%A8%A1%E5%9E%8B%E5%8F%91%E7%8E%B0671B%E6%A8%A1%E5%9E%8B%E7%9B%B2%E7%82%B9%23&amp;extparam=%23DeepSeek%E6%96%B07B%E6%A8%A1%E5%9E%8B%E5%8F%91%E7%8E%B0671B%E6%A8%A1%E5%9E%8B%E7%9B%B2%E7%82%B9%23" data-hide=""><span class="surl-text">#DeepSeek新7B模型发现671B模型盲点#</span></a><br><br>DeepSeek放大招！新模型专注数学定理证明，大幅刷新多项高难基准测试。<br><br>在普特南测试上，新模型DeepSeek-Prover-V2直接把记录刷新到49道。<br><br>目前的第一名在657道题中只做出10道题，为Kimi与AIME2024冠军团队Numina合作成果Kimina-Prover。<br><br>而未针对定理证明优化的DeepSeek-R1只做出1道。<br><br>让还没发布的R2更令人期待了。【图1】<br><br>除测评结果之外，论文中特别报告了“通过强化学习发现新技能”现象。<br><br>正如R1带来了“啊哈时刻”，Prover-V2也有令人意想不到的能力。【图2】<br><br>具体来说，在普特南测试中，参数量较小的DeepSeek-Prover-V2-7B用非CoT生成模式成功解决了13个671B模型未能解决的问题。<br><br>团队仔细检查该模型的输出后发现，其推理方法存在一个独特模式：7B模型处理涉及有限基数的问题时，经常使用Cardinal.toNat和Cardinal.natCast_inj，而671B模型生成的输出中明显没有这些内容。<br><br>要注意，7B模型是在DeepSeek-Prover-V1.5-Base模型基础上，先使用671B模型在强化学习阶段收集的数据微调，再执行强化学习得来的。<br><br>也就是说，7B模型学会了671B模型没有学会的新技能。【图3】<br><br>那么，DeepSeeK-Prover-V2如何炼成的呢？与前代相比又有哪些改进？请看：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FVO5PBVQQEeIBrzGFs9b6Ww" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">DeepSeek新数学模型刷爆记录！7B小模型自主发现671B模型不会的新技能</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i0zvcfsv9cj30wy0v6k3k.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0zvcpi35wj30zk0ccwmw.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0zvdcn9t8j30zk0lggr7.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

DeepSeek推出新数学证明模型DeepSeek-Prover-V2，在普特南测试中取得49道题的突破性成绩，远超当前最佳模型（10道）。特别值得注意的是，7B参数的小模型通过强化学习发现了671B大模型未掌握的解题方法，在有限基数问题上展现出独特推理模式。该模型基于V1.5-Base微调后经强化学习训练，不仅性能大幅提升，还自主发展出大模型不具备的新技能。这一突破展示了小模型通过特定训练可能超越超大模型的潜力，为AI数学推理领域带来新方向。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-03T03:17:41Z
- **目录日期**: 2025-05-03
