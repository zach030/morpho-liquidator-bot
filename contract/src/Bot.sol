// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "@src/IBot.sol";
import {SafeTransferLib} from "@solady/utils/SafeTransferLib.sol";

contract Bot is IBot{
    error NoProfit();
    error OnlyOwner();
    error OnlyMorpho();
    address public immutable owner;
    address public constant MORPHO = 0xBBBBBbbBBb9cC5e90e3b3Af64bdAF62C37EEFFCb;

    constructor() payable {
        owner = msg.sender;
    }

    receive() external payable {}

    // ------------modifiers----------------
    modifier onlyOwner() {
        if (msg.sender != owner) revert OnlyOwner();
        _;
    }

    modifier onlyMorpho() {
        if (msg.sender != address(MORPHO)) revert OnlyMorpho();
        _;
    }

    function withdrawERC20(address token, uint256 amount) external onlyOwner {
        SafeTransferLib.safeTransfer(token, msg.sender, amount);
    }

    function withdrawETH(uint256 amount) external onlyOwner {
        SafeTransferLib.safeTransferETH(msg.sender, amount);
    }

    function approveERC20(address token, address to, uint256 amount) external onlyOwner {
        SafeTransferLib.safeApprove(token, to, amount);
    }

    function morphoLiquidate(bytes32 id, address borrower, uint256 seizedAssets, address pair, bytes calldata swapData) external payable onlyOwner {
        IMorpho.MarketParams memory params = IMorpho(MORPHO).idToMarketParams(id);
        IMorpho(MORPHO).liquidate(
            params,
            borrower,
            seizedAssets,
            0,
            abi.encode(
                MorphoLiquidateData(
                    params.collateralToken,
                    params.loanToken,
                    seizedAssets,
                    pair,
                    swapData
                )
            )
        );
    }

    function onMorphoLiquidate(uint256 repaidAssets, bytes calldata data) external onlyMorpho {
        MorphoLiquidateData memory arb = abi.decode(data, (MorphoLiquidateData));
        (bool success,) = arb.pair.call(arb.swapData);
        if (!success) revert ("swap error");
        uint256 out = SafeTransferLib.balanceOf(arb.loanToken, address(this));
        if (out < repaidAssets) revert NoProfit();
        SafeTransferLib.safeApprove(arb.loanToken, MORPHO, repaidAssets);
    }
}